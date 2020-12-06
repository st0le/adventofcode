using System;
using System.Collections.Generic;
using System.Linq;

namespace _5
{
    class Program
    {
        static void Main(string[] args)
        {
            new Program().Go();
        }

        public void Go()
        {
            // Unit Test
            // Console.WriteLine(Decode("BFFFBBFRRR"));
            // Console.WriteLine(Decode("FFFBBBFRRR"));
            // Console.WriteLine(Decode("BBFFBBFRLL"));
            string line;
            int minId = int.MaxValue, maxId = int.MinValue;
            long sum = 0;
            while ((line = Console.ReadLine()) != null)
            {
                int id = Decode(line);
                sum += id;
                maxId = Math.Max(maxId, id);
                minId = Math.Min(minId, id);
            }
            Console.WriteLine(maxId);
            Console.WriteLine((((minId + maxId) * (maxId - minId + 1)) / 2) - sum);
        }

        private int Decode(string s)
        {
            int row = 0, col = 0;
            for (int i = 0; i < 7; i++)
            {
                if (s[i] == 'B')
                {
                    row |= 1;
                }
                row <<= 1;
            }
            row >>= 1;
            for (int i = 7; i < 10; i++)
            {
                if (s[i] == 'R')
                {
                    col |= 1;
                }
                col <<= 1;
            }
            col >>= 1;
            return (row << 3) | col;
        }
    }
}
