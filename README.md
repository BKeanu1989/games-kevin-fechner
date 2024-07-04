# Games from/for Kevin Fechner




## Ideas
- Test Deck
- Arschloch
- Mäxchen
- truth or drink
- wizard
- 1 - 10 / B / D / K / A - Pik(e), Herz, Karo,  

## Open / Closed Games
- this has two meanings
1. games which shouldn't be open to others // extra password protected
2. there is one idea, where games are "open to all, everyone can move any card", has a hand, able to take multiple cards from any deck 

## planned Features
- rooms with password
- SDK for opennes to other 
- PWA 
- grpc?


--
## Todos
- docker env var for own docker image ... 




## Stack
- Nginx with passwordprotected routes
- Go <!-- for now, elixir might be interesting? -->
    - Server
    - CLI
- sqlite
- wasm? 
- service worker / offline aka 
- web worker / websocket / open connection
- peer to peer network? / bei safeboxen angwendet
- ui
    - canvas for card placements 
        - listener for events / drop 
    - vue/react
    - htmx / templ, when server heavy only?

---
---



```txt
Peer to peer network
A peer-to-peer (P2P) network is a type of network where devices or computers are connected directly to each other, forming a network of equal nodes. In a P2P network, each node has the same level of authority and can act as both a client and a server, sharing resources and communicating directly with other nodes.

Key Characteristics

Decentralized: P2P networks do not rely on a central server or authority, instead, each node is responsible for its own data and communication.
Equal nodes: All nodes in the network have the same level of authority and can act as both clients and servers.
Direct communication: Nodes communicate directly with each other, without the need for a central server or intermediary.
Resource sharing: Nodes can share resources, such as files, bandwidth, or processing power, with other nodes in the network.
Examples

File sharing networks, where users can share files with each other directly.
Distributed computing networks, where nodes work together to solve complex problems or perform tasks.
Ad-hoc networks, where devices connect to each other temporarily to share resources or communicate.
Advantages

Scalability: P2P networks can scale more easily than traditional client-server networks, as new nodes can be added as needed.
Flexibility: P2P networks can be more flexible, as nodes can be added or removed as needed.
Security: P2P networks can be more secure, as there is no single point of failure or central authority that can be targeted.
Disadvantages

Complexity: P2P networks can be more complex to manage and maintain, as each node has its own responsibilities.
Security risks: P2P networks can be more vulnerable to security risks, as each node has access to the network and can potentially compromise the security of the network.
Quality of service: P2P networks can suffer from quality of service issues, as nodes may have varying levels of bandwidth, processing power, or other resources.
```


## Adhoc network via iot devices? / microcontroller

## Sources & Recommendations

In general, why always strive to be as open as possible and credits any authors. Please message me, if I forgot anyone/anything. 


## melkeydev/
https://docs.go-blueprint.dev/

```bash
go-blueprint create --name games-go-blueprint --framework chi --driver sqlite --advanced --feature websocket
```