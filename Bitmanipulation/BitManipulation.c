public class BitManipulation
{
    public static long GetBit(long n, int k)
    {
        return (n >> k) & 1L;
    }

    public static long SetBit(long n, int k)
    {
        return (1L << k) | n;
    }

    public static long ClearBit(long n, int k)
    {
        return n & ~(1L << k);
    }

    public static long ToggleBit(long n, int k)
    {
        return n ^ (1L << k);
    }

    public static bool IsPowerOfTwo(long n)
    {
        return n > 0 && (n & (n - 1)) == 0;
    }
}
