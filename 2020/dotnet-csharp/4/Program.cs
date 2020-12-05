using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Text.Json;
using System.Text.Json.Serialization;
using System.Text.RegularExpressions;

namespace _4
{
    class Program
    {
        private class Passport
        {
            public int byr { get; set; }
            public int iyr { get; set; }
            public int eyr { get; set; }
            public string hgt { get; set; }
            public string hcl { get; set; }
            public string ecl { get; set; }
            public string pid { get; set; }
            public string cid { get; set; }
            public Passport(string passport)
            {
                byr = ParseInt(passport, "byr");
                iyr = ParseInt(passport, "iyr");
                eyr = ParseInt(passport, "eyr");
                hgt = ParseString(passport, "hgt");
                hcl = ParseString(passport, "hcl");
                ecl = ParseString(passport, "ecl");
                pid = ParseString(passport, "pid");
                cid = ParseString(passport, "cid");
            }

            public bool IsValid1()
            {
                return byr > 0
                    && iyr > 0
                    && eyr > 0
                    && hgt != string.Empty
                    && hcl != string.Empty
                    && ecl != string.Empty
                    && pid != string.Empty;
            }

            public void Debug()
            {
                Console.WriteLine("byr " + (1920 <= byr && byr <= 2002) + " " + byr);
                Console.WriteLine("iyr " + (2010 <= iyr && iyr <= 2020) + " " + iyr);
                Console.WriteLine("eyr " + (2020 <= eyr && eyr <= 2030) + " " + eyr);
                Console.WriteLine("hgt " + IsValidHeight() + " " + hgt);
                Console.WriteLine("hcl " + Regex.IsMatch(hcl, "#[0-9a-f]{6}") + " " + hcl);
                Console.WriteLine("ecl " + Regex.IsMatch(ecl, "amb|blu|brn|gry|grn|hzl|oth") + " " + ecl);
                Console.WriteLine("pid " + Regex.IsMatch(pid, "[0-9]{9}") + " " + pid);
                Console.WriteLine();
            }

            public bool IsValid2()
            {
                return 1920 <= byr && byr <= 2002
                    && 2010 <= iyr && iyr <= 2020
                    && 2020 <= eyr && eyr <= 2030
                    && IsValidHeight()
                    && Regex.IsMatch(hcl, "^#[0-9a-f]{6}$")
                    && Regex.IsMatch(ecl, "^(amb|blu|brn|gry|grn|hzl|oth)$")
                    && Regex.IsMatch(pid, "^[0-9]{9}$");
            }

            private bool IsValidHeight()
            {
                Match m = Regex.Match(hgt, @"^(\d+)(in|cm)$");
                if (!m.Success)
                {
                    return false;
                }

                if (!int.TryParse(m.Groups[1].Value, out int v))
                {
                    return false;
                }

                switch (m.Groups[2].Value)
                {
                    case "cm":
                        return 150 <= v && v <= 193;
                    case "in":
                        return 59 <= v && v <= 76;
                }

                return false;
            }

            public int ParseInt(string passport, string key)
            {
                Match m = Regex.Match(passport, @$"{key}:(\d+)");
                if (m.Success)
                {
                    return int.Parse(m.Groups[1].Value);
                }

                return 0;
            }

            public string ParseString(string passport, string key)
            {
                Match m = Regex.Match(passport, @$"{key}:(#?\w+)");
                if (m.Success)
                {
                    return m.Groups[1].Value;
                }

                return string.Empty;
            }
        }

        static void Main(string[] args)
        {
            new Program().Go();
        }

        private void Go()
        {
            List<Passport> passports = new List<Passport>();
            string line, passport;
            StringBuilder sb = new StringBuilder();
            while ((line = Console.ReadLine()) != null)
            {
                if (string.IsNullOrWhiteSpace(line))
                {
                    passport = sb.ToString();
                    passports.Add(new Passport(passport));
                    sb.Clear();
                }
                else
                {
                    sb.Append(line).Append(' ');
                }
            }

            passport = sb.ToString();
            passports.Add(new Passport(passport));
            Console.WriteLine(passports.Count(p => p.IsValid1()));
            Console.WriteLine(passports.Count(p => p.IsValid1() && p.IsValid2()));
        }
    }
}
