class BitManipulation {
    static getBit(n, k) {
        return (n >> BigInt(k)) & 1n;
    }

    static setBit(n, k) {
        return (1n << BigInt(k)) | n;
    }

    static clearBit(n, k) {
        return n & ~(1n << BigInt(k));
    }

    static toggleBit(n, k) {
        return n ^ (1n << BigInt(k));
    }

    static isPowerOfTwo(n) {
        return n > 0n && (n & (n - 1n)) === 0n;
    }
}
