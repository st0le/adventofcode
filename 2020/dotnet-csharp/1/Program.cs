using System;
using System.Collections.Generic;
using System.IO;

namespace _1
{
    class Program
    {
        static void Main(string[] args)
        {
            int target = 2020;
            List<int> numbers = new();
            string line;
            while ((line = Console.ReadLine()) != null)
            {
                numbers.Add(int.Parse(line));
            }

            numbers.Sort();

            Console.WriteLine(solve2Sum(numbers, target));
            Console.WriteLine(solve3Sum(numbers, target));

        }

        private static int solve3Sum(List<int> numbers, int target)
        {
            for (int i = 0; i < numbers.Count; i++)
            {
                int targetLeft = target - numbers[i];
                int product = solve2Sum(numbers, i + 1, numbers.Count - 1, targetLeft);
                if (product != 0)
                {
                    return product * numbers[i];
                }
            }
            return 0;
        }

        private static int solve2Sum(List<int> numbers, int target)
        {
            return solve2Sum(numbers, 0, numbers.Count - 1, target);
        }

        private static int solve2Sum(List<int> numbers, int start, int end, int target)
        {
            int i = start, j = end;
            while (i < j)
            {
                int s = numbers[i] + numbers[j];
                if (s == target)
                {
                    return numbers[i] * numbers[j];
                }
                else
                if (s > target)
                {
                    j--;
                }
                else
                {
                    i++;
                }
            }
            return 0;
        }
    }
}
