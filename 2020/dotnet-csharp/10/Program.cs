using System;
using System.Collections.Generic;

namespace _10
{
    class Program
    {
        static void Main(string[] args)
        {
            List<int> arr = new List<int>();
            string line;
            while ((line = Console.ReadLine()) != null)
            {
                arr.Add(int.Parse(line));
            }

            // setup array with final volt
            int target;
            arr.Sort();
            arr.Add(target = arr[arr.Count - 1] + 3);

            // part 1
            Dictionary<int, int> freq = new Dictionary<int, int>();
            freq[1] = freq[2] = freq[3] = 0;
            int prev = 0;
            foreach (int v in arr)
            {
                freq[v - prev]++;
                prev = v;
            }
            Console.WriteLine(freq[1] * freq[3]);

            // part 2
            long[] ways = new long[target + 1];
            ways[0] = 1;
            foreach (int v in arr)
            {
                for (int k = 1; k <= 3; k++)
                {
                    if (v - k >= 0)
                    {
                        ways[v] += ways[v - k];
                    }
                }
            }
            Console.WriteLine(ways[target]);
        }
    }
}
