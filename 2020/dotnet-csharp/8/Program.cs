using System;
using System.Collections.Generic;
using System.Linq;
using System.Text.RegularExpressions;

namespace _8
{
    class Program
    {
        static void Main(string[] args)
        {
            List<string> lines = new List<string>();
            string line;
            while ((line = Console.ReadLine()) != null)
            {
                lines.Add(line);
            }

            Console.WriteLine(new Program().RunProgram(lines, out _));
            for (int i = 0; i < lines.Count; i++)
            {
                if (lines[i].StartsWith("nop"))
                {
                    string hold = lines[i];
                    lines[i] = lines[i].Replace("nop", "jmp");
                    int acc = new Program().RunProgram(lines, out bool terminated);
                    if (terminated)
                    {
                        Console.WriteLine(acc);
                    }
                    lines[i] = hold;
                }

                if (lines[i].StartsWith("jmp"))
                {
                    string hold = lines[i];
                    lines[i] = lines[i].Replace("jmp", "nop");
                    int acc = new Program().RunProgram(lines, out bool terminated);
                    if (terminated)
                    {
                        Console.WriteLine(acc);
                    }
                    lines[i] = hold;
                }
            }
        }

        private record OpCode(string op, int offset);

        private OpCode Parse(string opcode)
        {
            Match m = Regex.Match(opcode, @"(\w+) ([\+\-]\d+)");
            return new OpCode(m.Groups[1].Value, int.Parse(m.Groups[2].Value));
        }

        public int RunProgram(List<string> opcodes, out bool terminates)
        {
            OpCode[] code = opcodes.Select(o => Parse(o)).ToArray();
            bool[] executed = new bool[code.Length];
            int acc = 0;
            int ip = 0;
            while (ip < code.Length && !executed[ip])
            {
                // Console.WriteLine($"{ip} - {code[ip]} - {acc}");
                executed[ip] = true;
                switch (code[ip].op)
                {
                    case "jmp":
                        ip += code[ip].offset;
                        break;
                    case "nop":
                        ip++;
                        break;
                    case "acc":
                        acc += code[ip].offset;
                        ip++;
                        break;
                    default:
                        break;
                }
            }

            terminates = !(ip < code.Length);
            return acc;
        }
    }
}
