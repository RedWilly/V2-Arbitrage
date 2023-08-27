//SPDX-License-Identifier: Unlicense
pragma solidity ^0.7.0;

//import '@openzeppelin/contracts/token/ERC20/IERC20.sol';
//import '@openzeppelin/contracts/token/ERC20/SafeERC20.sol';
//import '@openzeppelin/contracts/access/Ownable.sol';
//import '@openzeppelin/contracts/utils/EnumerableSet.sol';
//import 'hardhat/console.sol';

import '@openzeppelin/contracts/token/ERC20/IERC20.sol';
import '@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol';
import '@openzeppelin/contracts/access/Ownable.sol';
import '@openzeppelin/contracts/utils/structs/EnumerableSet.sol';
import 'hardhat/console.sol';


import './interfaces/IUniswapV2Pair.sol';
import './interfaces/IWETH.sol';
import './libraries/Decimal.sol';
import "./libraries/SafeMath.sol";

// 两个池子的余额
struct OrderedReserves{
    // pool 1  reserve asset
    uint256 a1;
    uint256 b1;
    // pool 2 reserve asset
    uint256 a2;
    uint256 b2;
}

// 套利信息
struct ArbitrageInfo {
    // 基础代币
    address baseToken;
    // 用于交换的代币
    address quoteToken;
    // 确定基础代币的位置,
    //e.g  (token0, token1) 如果token1 是基础代币,那么是 false,
    //  token0 是基础代币，则 是 true
    bool baseTokenSmaller;
    address lowerPool;
    address higherPool;

}

// CallbackData 回调使用
struct CallbackData {
    // 在低价的池子中欠款的地址
    address debtPool;
    // 价格较高的池子
    address targetPool;
    bool debtTokenSmaller;
    // 借入的token
    address borrowedToken;
    // 欠款资产
    address debtToken;
    // 欠款数量
    uint256 debtAmount;
    // 在高价的池子中得到的欠款资产的数量
    uint256 debtTokenOutAmount;
}


contract FlashBot is Ownable{
    using Decimal for Decimal.D256;
    using SafeMath for uint256;
    using SafeERC20 for IERC20;
    using EnumerableSet for EnumerableSet.AddressSet;

    // ACCESS CONTROL
    // Only the `permissionedPairAddress` may call the `uniswapV2Call` function
    address permissionedPairAddress = address(1);


    // WETH on ETH or WBNB on BSC
    address immutable WETH;

    // AVAILABLE BASE TOKENS
    // 基础的代币集合
    EnumerableSet.AddressSet baseTokens;


    event Withdraw(address indexed to, uint256 indexed value);
    event BaseTokenAdded(address indexed token);
    event BaseTokenRemoved(address indexed token);

    // 初始化基础代币代币
    constructor(address _WETH) {
        WETH = _WETH;
        baseTokens.add(_WETH);
    }

    // 合约可以收款
    receive() external payable {}

    /// @dev Redirect uniswap callback function
    /// The callback function on different DEX are not same, so use a fallback to redirect to uniswapV2Call
    fallback(bytes calldata _input) external returns (bytes memory) {
        (address sender, uint256 amount0, uint256 amount1, bytes memory data) = abi.decode(_input[4:], (address, uint256, uint256, bytes));
        uniswapV2Call(sender, amount0, amount1, data);
    }

    // 提款
    function withdraw() external {
        // 本币资产
        uint256 balance = address(this).balance;
        if (balance > 0) {
            payable(owner()).transfer(balance);
            emit Withdraw(owner(), balance);
        }

        // 其他有价值的资产
        for (uint256 i = 0; i < baseTokens.length(); i++) {
            address token = baseTokens.at(i);
            balance = IERC20(token).balanceOf(address(this));
            if (balance > 0) {
                // do not use safe transfer here to prevents revert by any shitty token
                IERC20(token).transfer(owner(), balance);
            }
        }

    }

    // 添加有价值的资产
    function addBaseToken(address token) external onlyOwner {
        baseTokens.add(token);
        emit BaseTokenAdded(token);
    }

    // 移除某个资产
    function removeBaseToken(address token) external onlyOwner {
        uint256 balance = IERC20(token).balanceOf(address(this));
        if (balance > 0) {
            // do not use safe transfer to prevents revert by any shitty token
            IERC20(token).transfer(owner(), balance);
        }
        baseTokens.remove(token);
        emit BaseTokenRemoved(token);
    }


    function getBaseTokens() external view returns (address[] memory tokens) {
        uint256 length = baseTokens.length();
        tokens = new address[](length);
        for (uint256 i = 0; i < length; i++) {
            tokens[i] = baseTokens.at(i);
        }
    }

    // 某个代币是否在基础列表中
    function isExistOnBaseTokenSet(address token) public view returns(bool) {
        return baseTokens.contains(token);
    }

    // 基础资产在pair对的排序
    function isBaseTokenSmaller(address pool0, address pool1) internal view returns (bool baseTokenSmaller, address baseToken, address quoteToken){
        require(pool0 != pool1, "Same pair address");
        (address pool0Token0, address pool0Token1) = (IUniswapV2Pair(pool0).token0(), IUniswapV2Pair(pool0).token1());
        (address pool1Token0, address pool1Token1) = (IUniswapV2Pair(pool1).token0(), IUniswapV2Pair(pool1).token1());

        // 检查是否是标准的 uniSwapV2 币对
        require(pool0Token0 < pool0Token1 && pool1Token0 < pool1Token1, 'Non standard uniswap AMM pair');
        // 需要 币对是一样的
        require(pool0Token0 == pool1Token0 && pool0Token1 == pool1Token1, 'Require same token pair');
        // 是否有我们需要的 代币
        require(isExistOnBaseTokenSet(pool0Token0) || isExistOnBaseTokenSet(pool0Token1), 'No base token in pair');

        // 判断 token0 是否是 基础代币
        (baseTokenSmaller, baseToken, quoteToken) =  isExistOnBaseTokenSet(pool0Token0)
            ? (true, pool0Token0, pool0Token1) : (false,pool0Token1, pool0Token0);
    }

    /// @dev Compare price denominated in quote token between two pools
    /// We borrow base token by using flash swap from lower price pool and sell them to higher price pool
    // 使用 quote 代币报价
    function getOrderedReserves(address pool0, address pool1, bool baseTokenSmaller) internal view returns(address lowerPool, address higherPool,OrderedReserves memory orderedReserves ){

        // 池子的流动性，还剩有多少代币
        (uint256 pool0Reserve0, uint256 pool0Reserve1 ) = IUniswapV2Pair(pool0).getReserves();
        (uint256 pool1Reserve0, uint256 pool1Reserve1 ) = IUniswapV2Pair(pool1).getReserves();

        // 计算价格，使用报价代币计算
        // token0 是基础代币, token1(quote) 的价格为： price = rToken0 / rToken1
        // token1 是基础代币， token0(quote) 的价格为： price = rToken1 / rToken0
        (Decimal.D256 memory price0, Decimal.D256 memory price0)=
            baseTokenSmaller
                ?(Decimal.from(pool0Reserve0).div(pool0Reserve1), Decimal.from(pool1Reserve0).div(pool1Reserve1))
                :(Decimal.from(pool0Reserve1).div(pool0Reserve0), Decimal.from(pool1Reserve1).div(pool1Reserve0));

        // get a1, b1, a2, b2 with following rule:
        // 1. (a1, b1) represents the pool with lower price, denominated in quote asset token
        // 2. (a1, a2) are the base tokens in two pools
        if (price0.lessThan(price1)) {
            (lowerPool, higherPool) = (pool0, pool1);
            (orderedReserves.a1, orderedReserves.b1, orderedReserves.a2, orderedReserves.b2) = baseTokenSmaller
                ? (pool0Reserve0,pool0Reserve1,pool1Reserve0, pool1Reserve1)
                : (pool0Reserve1, pool0Reserve0,pool1Reserve1,  pool1Reserve0);
        } else{
            (lowerPool, higherPool) = (pool1, pool0);
            (orderedReserves.a1, orderedReserves.b1, orderedReserves.a2, orderedReserves.b2) = baseTokenSmaller
            ? (pool1Reserve0, pool1Reserve1,pool0Reserve0, pool0Reserve1)
            : (pool1Reserve1, pool1Reserve0,pool0Reserve1,pool0Reserve0);
        }
        console.log('Borrow from pool:', lowerPool);
        console.log('Sell to pool:', higherPool);
    }

    /// @notice Calculate how much profit we can by arbitraging between two pools
    // 计算利润
    function getProfit(address pool0, address pool1) external view returns (uint256 profit, address baseToken) {
        (bool baseTokenSmaller, , ) = isBaseTokenSmaller(pool0, pool1);
        baseToken = baseTokenSmaller ? IUniswapV2Pair(pool0).token0() : IUniswapV2Pair(pool0).token1();

        (, , OrderedReserves memory orderedReserves) = getOrderedReserves(pool0, pool1, baseTokenSmaller);

        uint256 borrowAmount = calcBorrowAmount(orderedReserves);
        // borrow quote token on lower price pool,
        // borrowAmount 是从低价池子中借出 quote token数量
        // getAmountIn 计算 借出了 quote token borrowAmount， 需要 打入 base token 的数量
        uint256 debtAmount = getAmountIn(borrowAmount, orderedReserves.a1, orderedReserves.b1);
        // sell borrowed quote token on higher price pool
        // getAmountOut 打入  quote token borrowAmount 个， 可以换出 base token 的数量
        uint256 baseTokenOutAmount = getAmountOut(borrowAmount, orderedReserves.b2, orderedReserves.a2);
        if (baseTokenOutAmount < debtAmount) {
            profit = 0;
        } else {
            profit = baseTokenOutAmount - debtAmount;
        }
    }

    /// @notice Do an arbitrage between two Uniswap-like AMM pools
    /// @dev Two pools must contains same token pair
    function flashArbitrage(address pool0, address pool1) external {
        ArbitrageInfo memory info;
        // 确定基础代币的的位置
        (info.baseTokenSmaller, info.baseToken, info.quoteToken) = isBaseTokenSmaller(pool0, pool1);

        // 确定价格
        OrderedReserves memory orderedReserves;
        (info.lowerPool, info.higherPool, orderedReserves) = getOrderedReserves(pool0, pool1, info.baseTokenSmaller);

        // this must be updated every transaction for callback origin authentication
        permissionedPairAddress = info.lowerPool;


        // 套利前的代币余额
        uint256 balanceBefore = IERC20(info.baseToken).balanceOf(address(this));


        // avoid stack too deep error
        {
            // 计算 quote token 借出的数量
            uint256 borrowAmount = calcBorrowAmount(orderedReserves);
            (uint256 amount0Out, uint256 amount1Out) = info.baseTokenSmaller ? (uint256(0), borrowAmount) : (borrowAmount, uint256(0));

            // borrow quote token on lower price pool, calculate how much debt we need to pay demoninated in base token
            // getAmountIn 计算 借出了 quote token borrowAmount， 需要 打入 base token 的数量
            uint256 debtAmount = getAmountIn(borrowAmount, orderedReserves.a1, orderedReserves.b1);

            // sell borrowed quote token on higher price pool, calculate how much base token we can get
            // getAmountOut 打入  quote token borrowAmount 个， 可以换出 base token 的数量
            uint256 baseTokenOutAmount = getAmountOut(borrowAmount, orderedReserves.b2, orderedReserves.a2);

            // 确保有利可图
            require(baseTokenOutAmount > debtAmount, 'Arbitrage fail, no profit');
            console.log('Profit:', (baseTokenOutAmount - debtAmount) / 1 ether);

            // can only initialize this way to avoid stack too deep error
            CallbackData memory callbackData;
            callbackData.debtPool = info.lowerPool;
            callbackData.targetPool = info.higherPool;
            callbackData.debtTokenSmaller = info.baseTokenSmaller;
            callbackData.borrowedToken = info.quoteToken;
            callbackData.debtToken = info.baseToken;
            callbackData.debtAmount = debtAmount;
            callbackData.debtTokenOutAmount = baseTokenOutAmount;

            bytes memory data = abi.encode(callbackData);

            IUniswapV2Pair(info.lowerPool).swap(amount0Out, amount1Out, address(this), data);
        }

        uint256 balanceAfter = IERC20(info.baseToken).balanceOf(address(this));
        require(balanceAfter > balanceBefore, 'Losing money');

        if (info.baseToken == WETH) {
            IWETH(info.baseToken).withdraw(balanceAfter);
        }
        permissionedPairAddress = address(1);

    }


    function uniswapV2Call(
        address sender,
        uint256 amount0,
        uint256 amount1,
        bytes memory data
    ) public {
        // access control
        require(msg.sender == permissionedPairAddress, 'Non permissioned address call');
        require(sender == address(this), 'Not from this contract');

        uint256 borrowedAmount = amount0 > 0 ? amount0 : amount1;
        CallbackData memory info = abi.decode(data, (CallbackData));

        IERC20(info.borrowedToken).safeTransfer(info.targetPool, borrowedAmount);

        (uint256 amount0Out, uint256 amount1Out) =
        info.debtTokenSmaller ? (info.debtTokenOutAmount, uint256(0)) : (uint256(0), info.debtTokenOutAmount);
        IUniswapV2Pair(info.targetPool).swap(amount0Out, amount1Out, address(this), new bytes(0));

        IERC20(info.debtToken).safeTransfer(info.debtPool, info.debtAmount);
    }

    // ***************
    //    计算部分
    // ***************

    /// @dev calculate the maximum base asset amount to borrow in order to get maximum profit during arbitrage
    function calcBorrowAmount(OrderedReserves memory reserves) internal pure returns (uint256 amount) {
        // we can't use a1,b1,a2,b2 directly, because it will result overflow/underflow on the intermediate result
        // so we:
        //    1. divide all the numbers by d to prevent from overflow/underflow
        //    2. calculate the result by using above numbers
        //    3. multiply d with the result to get the final result
        // Note: this workaround is only suitable for ERC20 token with 18 decimals, which I believe most tokens do

        uint256 min1 = reserves.a1 < reserves.b1 ? reserves.a1 : reserves.b1;
        uint256 min2 = reserves.a2 < reserves.b2 ? reserves.a2 : reserves.b2;
        uint256 min = min1 < min2 ? min1 : min2;

        // choose appropriate number to divide based on the minimum number
        uint256 d;
        if (min > 1e24) {
            d = 1e20;
        } else if (min > 1e23) {
            d = 1e19;
        } else if (min > 1e22) {
            d = 1e18;
        } else if (min > 1e21) {
            d = 1e17;
        } else if (min > 1e20) {
            d = 1e16;
        } else if (min > 1e19) {
            d = 1e15;
        } else if (min > 1e18) {
            d = 1e14;
        } else if (min > 1e17) {
            d = 1e13;
        } else if (min > 1e16) {
            d = 1e12;
        } else if (min > 1e15) {
            d = 1e11;
        } else {
            d = 1e10;
        }

        (int256 a1, int256 a2, int256 b1, int256 b2) =
        (int256(reserves.a1 / d), int256(reserves.a2 / d), int256(reserves.b1 / d), int256(reserves.b2 / d));

        int256 a = a1 * b1 - a2 * b2;
        int256 b = 2 * b1 * b2 * (a1 + a2);
        int256 c = b1 * b2 * (a1 * b2 - a2 * b1);

        (int256 x1, int256 x2) = calcSolutionForQuadratic(a, b, c);

        // 0 < x < b1 and 0 < x < b2
        require((x1 > 0 && x1 < b1 && x1 < b2) || (x2 > 0 && x2 < b1 && x2 < b2), 'Wrong input order');
        amount = (x1 > 0 && x1 < b1 && x1 < b2) ? uint256(x1) * d : uint256(x2) * d;
    }

    /// @dev find solution of quadratic equation: ax^2 + bx + c = 0, only return the positive solution
    function calcSolutionForQuadratic(
        int256 a,
        int256 b,
        int256 c
    ) internal pure returns (int256 x1, int256 x2) {
        int256 m = b**2 - 4 * a * c;
        // m < 0 leads to complex number
        require(m > 0, 'Complex number');

        int256 sqrtM = int256(sqrt(uint256(m)));
        x1 = (-b + sqrtM) / (2 * a);
        x2 = (-b - sqrtM) / (2 * a);
    }

    /// @dev Newton’s method for caculating square root of n
    function sqrt(uint256 n) internal pure returns (uint256 res) {
        assert(n > 1);

        // The scale factor is a crude way to turn everything into integer calcs.
        // Actually do (n * 10 ^ 4) ^ (1/2)
        uint256 _n = n * 10**6;
        uint256 c = _n;
        res = _n;

        uint256 xi;
        while (true) {
            xi = (res + c / res) / 2;
            // don't need be too precise to save gas
            if (res - xi < 1000) {
                break;
            }
            res = xi;
        }
        res = res / 10**3;
    }

    // copy from UniswapV2Library
    // given an output amount of an asset and pair reserves, returns a required input amount of the other asset
    function getAmountIn(
        uint256 amountOut,
        uint256 reserveIn,
        uint256 reserveOut
    ) internal pure returns (uint256 amountIn) {
        require(amountOut > 0, 'UniswapV2Library: INSUFFICIENT_OUTPUT_AMOUNT');
        require(reserveIn > 0 && reserveOut > 0, 'UniswapV2Library: INSUFFICIENT_LIQUIDITY');
        uint256 numerator = reserveIn.mul(amountOut).mul(1000);
        uint256 denominator = reserveOut.sub(amountOut).mul(997);
        amountIn = (numerator / denominator).add(1);
    }

    // copy from UniswapV2Library
    // given an input amount of an asset and pair reserves, returns the maximum output amount of the other asset
    function getAmountOut(
        uint256 amountIn,
        uint256 reserveIn,
        uint256 reserveOut
    ) internal pure returns (uint256 amountOut) {
        require(amountIn > 0, 'UniswapV2Library: INSUFFICIENT_INPUT_AMOUNT');
        require(reserveIn > 0 && reserveOut > 0, 'UniswapV2Library: INSUFFICIENT_LIQUIDITY');
        uint256 amountInWithFee = amountIn.mul(997);
        uint256 numerator = amountInWithFee.mul(reserveOut);
        uint256 denominator = reserveIn.mul(1000).add(amountInWithFee);
        amountOut = numerator / denominator;
    }
}
