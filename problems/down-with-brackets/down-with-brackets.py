from turtle import down
from typing import Callable


def down_with_brackets(s: str) -> str:
    """Returns YES if bracket string can be broken, NO if not;


    Bracket string can be broken iff contains the characters ')(' next to each other (maybe)
    """

    for c, cn in zip(s, s[1:]):
        if c == ")" and cn == "(":
            return "YES"

    return "NO"


def debug_parser(fn: Callable) -> None:
    tests = int(input())
    print(f"{tests} test cases")

    for i in range(tests):
        test = input()
        print(f"{i:>3}\t{test[:20]:<20}\t{down_with_brackets(test):>3}")


def comp_parser(fn: Callable) -> None:
    tests = int(input())

    for i in range(tests):
        print(fn(input()))


if __name__ == "__main__":
    comp_parser(down_with_brackets)
