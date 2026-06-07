using System;

namespace CpuSchedulingDotNet.Algorithms
{
  public static class LongestJobFirst
  {
    public static SchedulingResult Execute(Process[] processes)
    {
      var sortedProcesses = processes.OrderBy(p => p.BurstTime).Reverse().ToArray();
      Console.WriteLine(sortedProcesses.ToString());
      double totalTurnaroundTime = 0;
      double totalWaitTime = 0;
      double currTime = 0;

      foreach (var process in sortedProcesses)
      {
        double completionTime = Math.Max(currTime, process.ArrivalTime) + process.BurstTime;
        double turnaroundTime = completionTime - process.ArrivalTime;
        double waitTime = turnaroundTime - process.BurstTime;

        totalWaitTime += waitTime;
        totalTurnaroundTime += turnaroundTime;
        currTime = completionTime;
      }

      return new SchedulingResult(totalWaitTime / sortedProcesses.Length, totalTurnaroundTime / sortedProcesses.Length);
    }
  }
}