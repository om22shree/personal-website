from typing import List


# Pattern: bounded Bellman-Ford over edge count.
# Invariant: after i rounds, prices uses at most i edges.
# Complexity: O(k * E) time, O(V) space.
# Interview line: copy the previous price array so one round only adds one extra flight.
def cheapest_flight_with_k_stops(
    n: int, flights: List[List[int]], src: int, dst: int, k: int
) -> int:
    prices = [float("inf")] * n
    prices[src] = 0

    for _ in range(k + 1):
        next_prices = prices[:]
        for a, b, price in flights:
            if prices[a] != float("inf"):
                next_prices[b] = min(next_prices[b], prices[a] + price)
        prices = next_prices

    return -1 if prices[dst] == float("inf") else int(prices[dst])
