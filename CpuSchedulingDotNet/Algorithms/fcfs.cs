using System;
using System.Linq;

namespace CpuSchedulingDotNet.Algorithms;

/// <summary>
/// Implements first come first server CPU scheduling algorithm
/// </summary>
public static class FirstComeFirstServeAlgorithm
{
  public static SchedulingResult Execute(Process[] processes)
  {
    ArgumentNullException.ThrowIfNull(processes);

    if (processes.Length == 0)
      return new SchedulingResult(0, 0);

    var sortedProcesses = processes.OrderBy(p => p.ArrivalTime).ToArray();

    int processCount = sortedProcesses.Length;
    double currTime = 0;
    double totalWaitTime = 0;
    double totalTurnaroundTime = 0;

    foreach (var process in sortedProcesses)
    {
      double completionTime = Math.Max(currTime, process.ArrivalTime) + process.BurstTime;
      double turnaroundTime = completionTime - process.ArrivalTime;
      double waitTime = turnaroundTime - process.BurstTime;

      totalWaitTime += waitTime;
      totalTurnaroundTime += turnaroundTime;

      currTime = completionTime;
    }

    double averageWaitingTime = totalWaitTime / processCount;
    double averageTurnaroundTime = totalTurnaroundTime / processCount;

    return new SchedulingResult(averageWaitingTime, averageTurnaroundTime);
  }
}
