# Adaptive Consensus

<div style="text-align: justify">
Consensus is essential to the blockchain as it enables participants to share a consistent view of the underlying distributed ledger. Currently existing protocols either rely on Proof-of-Work (PoW) or similar economic incentive schemes, with high transaction latency but that can handle thousands of participants or on classical byzantine fault tolerant consensus protocols, with low transaction latency but that do not scale well with the number of participants. 

The Adaptive Consensus is a proposal of an consensus protocol, in order to improve the scalability. The procotol use an optimization mechanism to
nd the ideal parameter values, i.e., different settings (e.g. network), scale (participants), load, etc. 

It is based on the paper [Worldwide Consensus](https://www.researchgate.net/publication/220973671_Worldwide_Consensus).
</div>

## Run

To run a peer, only need to do the following steps:

    make
    cd ./consensus/bin/
    peer <YOUR_PORT> <NEIGHTBOR_PORT> ...


