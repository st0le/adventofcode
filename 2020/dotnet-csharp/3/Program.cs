using System;
using System.Collections.Generic;

namespace _3
{
    class Program
    {
        static void Main(string[] args)
        {
            List<string> lines = new();
            string line;
            while ((line = Console.ReadLine()) != null)
            {
                lines.Add(line);
            }

            // part 1
            Console.WriteLine(CalculateTrees(lines, 3, 1));

            // part 2
            int[][] slopes = new int[][]{
                new int[]{1,1},
                new int[]{3,1},
                new int[]{5,1},
                new int[]{7,1},
                new int[]{1,2}
            };
            int product = 1;
            foreach (var slope in slopes)
            {
                product *= CalculateTrees(lines, slope[0], slope[1]);
            }
            Console.WriteLine(product);
        }

        static int CalculateTrees(IList<string> lines, int dx, int dy)
        {
            int count = 0;
            int rows = lines.Count, cols = lines[0].Length;
            for (int x = 0, y = 0; y < rows; y += dy, x += dx)
            {
                if (lines[y][x % cols] == '#')
                {
                    count++;
                }
            }
            return count;
        }
    }
}
