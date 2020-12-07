using System;
using System.Collections.Generic;
using System.Linq;
using System.Text.RegularExpressions;

namespace _7
{
    class Program
    {
        static void Main(string[] args)
        {
            Dictionary<string, List<Tuple<string, int>>> rules = new();
            string line;
            while (null != (line = Console.ReadLine()))
            {
                string container = line.Substring(0, line.IndexOf(" bags"));
                var list = new List<Tuple<string, int>>();
                rules.Add(container, list);
                foreach (Match m in Regex.Matches(line, @"(\d+) (\w+\s\w+)"))
                {
                    int count = int.Parse(m.Groups[1].Value);
                    string smallerBag = m.Groups[2].Value;
                    list.Add(Tuple.Create(smallerBag, count));
                }
            }


            // foreach (var item in rules)
            // {
            //     Console.WriteLine(item.Key);
            // }
            Console.WriteLine(rules.Keys.Where(bag => bag != "shiny gold" && CanContain(rules, bag, "shiny gold")).Count());
            Console.WriteLine(CalculateBags(rules, "shiny gold") - 1);
        }

        static bool CanContain(Dictionary<string, List<Tuple<string, int>>> rules, string sourceBag, string targetBag)
        {
            if (sourceBag == targetBag)
            {
                return true;
            }

            foreach (var item in rules[sourceBag])
            {
                if (CanContain(rules, item.Item1, targetBag))
                {
                    return true;
                }
            }
            return false;
        }

        static int CalculateBags(Dictionary<string, List<Tuple<string, int>>> rules, string bag)
        {
            int s = 1;
            foreach (var item in rules[bag])
            {
                s += item.Item2 * CalculateBags(rules, item.Item1);
            }

            return s;
        }
    }
}
