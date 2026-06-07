using System;
using System.Linq;

namespace CpuSchedulingDotNet.Algorithms;

/// <summary>
/// Implements first come first server CPU scheduling algorithm
/// </summary>
public class FirstComeFirstServeAlgorithm
{

  public static SchedulingResult Execute(Process[] processes)
  {

    ArgumentNullException.ThrowIfNull(processes);

    if (processes.Length == 0)
      return new SchedulingResult(0, 0);

    Process[] sortedProcesses = [.. processes.OrderBy(p => p.ArrivalTime)];

    int processCount = sortedProcesses.Length;
    double[] completionTime = new double[processCount];
    double[] turnAroundTime = new double[processCount];

    completionTime[0] = sortedProcesses[0].ArrivalTime + sortedProcesses[0].BurstTime;
    turnAroundTime[0] = processes[0].BurstTime;

    double totalWaitTime = 0;
    double totalTurnaroundTime = turnAroundTime[0];

    for (int i = 1; i < processCount; i++)
    {
      Process currProcess = sortedProcesses[i];

      completionTime[i] = Math.Max(currProcess.ArrivalTime, completionTime[i - 1]) + currProcess.BurstTime;
      turnAroundTime[i] = completionTime[i] - currProcess.ArrivalTime;

      totalWaitTime += turnAroundTime[i] - currProcess.BurstTime;
      totalTurnaroundTime += turnAroundTime[i];
    }

    double averageWaitingTime = totalWaitTime / processCount;
    double averageTurnAroundTime = totalTurnaroundTime / processCount;

    Console.WriteLine("Average Wait Time: " + averageWaitingTime);
    Console.WriteLine("Average Turn Around Time: " + averageTurnAroundTime);

    return new SchedulingResult(averageWaitingTime, averageTurnAroundTime);
  }
}
