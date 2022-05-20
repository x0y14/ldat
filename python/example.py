class Name:
    def __init__(self, first: str, middle: str, last: str):
        self.first: str = first
        self.middle: str = middle
        self.last: str = last

    def to_s(self) -> None:
        print(self.first + " " + self.middle + " " + self.last)
