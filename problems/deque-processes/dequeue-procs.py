#!/usr/bin/python3


def dequeue_procs(permutation: list[int], len_p: int) -> str:
    l, r = 0, len_p - 1
    take_min = True

    results = ["L"] * len_p
    i = 0

    while l != r:
        if take_min:
            if permutation[l] < permutation[r]:
                l += 1
            else:
                r -= 1
                results[i] = "R"
        else:
            if permutation[r] > permutation[l]:
                r -= 1
                results[i] = "R"
            else:
                l += 1

        i += 1
        take_min = not take_min

    return "".join(results)


if __name__ == "__main__":
    testcases = int(input())

    for _ in range(testcases):
        len_p = int(input())
        permutation = [int(p) for p in input().split()]

        print(dequeue_procs(permutation, len_p))
