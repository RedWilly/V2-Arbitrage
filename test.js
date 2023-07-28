const axios = require('axios');

/*
testing rpc letancy/speed and responds
*/

async function testRPC(url) {
    try {
        const startTime = Date.now();
        await axios.get(url);
        const endTime = Date.now();
        const latency = endTime - startTime;
        return latency;
    } catch (error) {
        console.error(`Error testing RPC at ${url}:`, error.message);
        return null;
    }
}

async function main() {
    const rpcEndpoints = [
        {
            name: 'Ankr',
            url: `https://rpc.ankr.com/dogechain`,
        },
        {
            name: 'doge1',
            url: `https://rpc-us.dogechain.dog`,
        },
        {
            name: 'doge2',
            url: `https://rpc.dogechain.dog	`,
        },
    ];

    const results = await Promise.all(
        rpcEndpoints.map(async (rpc) => {
            const latency = await testRPC(rpc.url);
            return {
                name: rpc.name,
                latency: latency !== null ? latency : Infinity,
            };
        })
    );

    results.sort((a, b) => a.latency - b.latency);

    console.log('Latency results:');
    results.forEach((result) => {
        console.log(`${result.name}: ${result.latency === Infinity ? 'Failed to test' : result.latency + ' ms'}`);
    });

    if (results.length >= 2) {
        const fastestRPC = results[0].name;
        const slowestRPC = results[results.length - 1].name;
        console.log(`\nFastest RPC: ${fastestRPC}`);
        console.log(`Slowest RPC: ${slowestRPC}`);
    }
}

main();
