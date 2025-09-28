import time


def prefix_min_suffix_max(a: list[int], len_a: int) -> str:
    res = ""

    for i, e in enumerate(a):
        e_1 = 0
        e_0 = float("inf")

        # TODO: make efficient with sliding window for mins,maxes
        if a[:i]:
            e_0 = min(a[:i])
        if a[i + 1 :]:
            e_1 = max(a[i + 1 :])

        if e_0 < e < e_1:
            res += "0"
            continue

        res += "1"

    return res


def prefix_min_suffix_max_opt(a: list[int], len_a: int) -> str:
    # Optimisation: pre-compute mins, maxes for each index of the array
    precom_mins: list[int] = [0] * len_a
    precom_mins[0] = a[0]

    precom_maxes: list[int] = [0] * len_a
    precom_maxes[-1] = a[-1]

    # quicker than str building (fewer allocations)
    res: list[str] = ["1"] * len_a

    # condition is 0 iff e_0 < e < e_1;
    # if we don't have e_0 or e_1, cannot be "0"
    if len_a <= 2:
        return "1" * len_a

    for i in range(1, len_a):
        precom_mins[i] = min(precom_mins[i - 1], a[i])

    for i in range(len_a - 2, -1, -1):
        precom_maxes[i] = max(precom_maxes[i + 1], a[i])

    for i in range(1, len_a - 1):
        e_0 = precom_mins[i - 1]
        e = a[i]
        e_1 = precom_maxes[i + 1]

        if e_0 < e < e_1:
            res[i] = "0"

    # also, know outers will always be viable;
    # either it is less than the max of what follows it, or
    # its greater than the min of what preceeds it
    return "".join(res)


def compare():
    testcases = int(input())

    for _ in range(testcases):
        len_a = int(input())
        a = [int(e) for e in input().split()]

        unop_start = time.time()
        unop = prefix_min_suffix_max(a, len_a)
        unop_end = time.time()

        op_start = time.time()
        op = prefix_min_suffix_max_opt(a, len_a)
        op_end = time.time()

        if op != unop:
            print("FAILED")
            print(f"\tEXPECTED {unop}\n\tGOT {op}")
        else:
            print(
                f"PASSED: unop took {unop_end - unop_start} vs op {op_end - unop_start}"
            )


if __name__ == "__main__":
    testcases = int(input())

    for _ in range(testcases):
        len_a = int(input())
        a = [int(e) for e in input().split()]

        print(prefix_min_suffix_max_opt(a, len_a))
