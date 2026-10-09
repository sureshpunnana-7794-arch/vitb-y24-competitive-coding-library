class BitManipulation:
    @staticmethod
    def get_bit(n: int, k: int) -> int:
        return (n >> k) & 1

    @staticmethod
    def set_bit(n: int, k: int) -> int:
        return (1 << k) | n

    @staticmethod
    def clear_bit(n: int, k: int) -> int:
        return n & ~(1 << k)

    @staticmethod
    def toggle_bit(n: int, k: int) -> int:
        return n ^ (1 << k)

    @staticmethod
    def is_power_of_two(n: int) -> bool:
        return n > 0 and (n & (n - 1)) == 0
