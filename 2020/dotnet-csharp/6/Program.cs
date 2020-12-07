using System;
using System.Collections.Generic;

namespace _6
{
    class Program
    {
        static void Main(string[] args)
        {
            int part1 = 0, part2 = 0, people = 0;
            Dictionary<char, int> answers = new Dictionary<char, int>();
            string line;
            while ((line = Console.ReadLine()) != null)
            {

                if (string.IsNullOrEmpty(line))
                {
                    // group ends
                    part1 += answers.Count;

                    foreach (var kv in answers)
                    {
                        if (kv.Value == people)
                        {
                            part2++;
                        }
                    }

                    answers.Clear();
                    people = 0;
                }
                else
                {
                    people++;
                    foreach (char c in line)
                    {
                        if (!answers.TryGetValue(c, out int v))
                        {
                            answers[c] = 1;
                        }
                        else
                        {
                            answers[c] = v + 1;
                        }
                    }
                }
            }

            // group ends
            part1 += answers.Count;
            foreach (var kv in answers)
            {
                if (kv.Value == people)
                {
                    part2++;
                }
            }
            answers.Clear();

            Console.WriteLine(part1);
            Console.WriteLine(part2);
        }
    }
}
