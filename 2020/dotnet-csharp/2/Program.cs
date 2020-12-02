using System;
using System.Linq;

namespace _2
{
    class Program
    {
        static void Main(string[] args)
        {
            string line;
            int firstPolicyValidCount = 0, secondPolicyValidCount = 0;
            while ((line = Console.ReadLine()) != null)
            {
                string[] parts = line.Split(" ");
                string[] lens = parts[0].Split("-");
                int low = int.Parse(lens[0]);
                int high = int.Parse(lens[1]);
                char c = parts[1][0];
                int count = parts[2].Count(ch => c == ch);
                if (low <= count && count <= high)
                {
                    firstPolicyValidCount++;
                }

                if (parts[2][low - 1] == c ^ parts[2][high - 1] == c)
                {
                    secondPolicyValidCount++;
                }

            }
            Console.WriteLine(firstPolicyValidCount);
            Console.WriteLine(secondPolicyValidCount);
        }
    }
}
