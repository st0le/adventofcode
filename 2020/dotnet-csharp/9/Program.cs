using System;
using System.Collections.Generic;

namespace _9
{
    class Program
    {
        static void Main(string[] args)
        {
            List<long> nums = new List<long>();
            string line;
            while ((line = Console.ReadLine()) != null)
            {
                nums.Add(long.Parse(line));
            }

            int N = nums.Count < 25 ? 5 : 25;
            long first = 0;
            for (int i = N; i < nums.Count; i++)
            {
                if (!CanAdd(nums, i - N, i, nums[i]))
                {
                    first = nums[i];
                    break;
                }
            }

            Console.WriteLine(first);
            Console.WriteLine(FindWeak(nums, first));

        }

        private static long FindWeak(List<long> nums, long target)
        {
            for (int i = 0; i < nums.Count; i++)
            {
                long min = nums[i], max = nums[i];
                long s = nums[i];
                for (int j = i + 1; j < nums.Count && s < target; j++)
                {
                    min = Math.Min(min, nums[j]);
                    max = Math.Max(max, nums[j]);
                    s += nums[j];
                    if (s == target)
                    {
                        return min + max;
                    }
                }
            }
            return -1;
        }

        private static bool CanAdd(List<long> nums, int lo, int hi, long target)
        {
            for (int i = lo; i < hi; i++)
            {
                for (int j = i + 1; j < hi; j++)
                {
                    if (nums[i] + nums[j] == target)
                    {
                        return true;
                    }
                }
            }
            return false;
        }
    }
}
