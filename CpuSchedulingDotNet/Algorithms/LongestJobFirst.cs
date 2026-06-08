using System;
using System.Linq;

namespace CpuSchedulingDotNet.Algorithms
{
  public static class LongestJobFirst
  {
    public static SchedulingResult Execute(Process[] processes)
    {
      var sortedProcesses = processes.OrderBy(p => p.ArrivalTime).ToArray();
      bool[] completed = new bool[sortedProcesses.Length];

      double totalTurnaroundTime = 0;
      double totalWaitTime = 0;
      double currTime = 0;

      int completedCount = 0;
      int largestBurstIdx = 0;

      while (completedCount < processes.Length)
      {
        largestBurstIdx = -1;
        for (int i = 0; i < sortedProcesses.Length; i++)
        {
          var process = sortedProcesses[i];

          if (completed[i])
            continue;

          if (
            process.ArrivalTime <= currTime &&
            (largestBurstIdx == -1 || process.BurstTime > sortedProcesses[largestBurstIdx].BurstTime)
          )
            largestBurstIdx = i;

        }

        if (largestBurstIdx == -1)
        {
          currTime++;
          continue;
        }

        var currProcess = sortedProcesses[largestBurstIdx];
        double completionTime = currTime + currProcess.BurstTime;
        double turnaroundTime = completionTime - currProcess.ArrivalTime;
        double waitTime = turnaroundTime - currProcess.BurstTime;

        totalTurnaroundTime += turnaroundTime;
        totalWaitTime += waitTime;
        currTime = completionTime;
        completed[largestBurstIdx] = true;
        completedCount++;
      }

      return new SchedulingResult(totalWaitTime / sortedProcesses.Length, totalTurnaroundTime / sortedProcesses.Length);
    }
  }
}
