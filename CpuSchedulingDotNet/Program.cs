using CpuSchedulingDotNet.Algorithms;

namespace CpuSchedulingDotNet
{
  public struct Process
  {
    public int Priority;
    public int ArrivalTime;
    public int BurstTime;
  };

  public record SchedulingResult(
    double AverageTurnaroundTime,
    double AverageWaitTime
  );

  public class Runner
  {
    public static void Main(string[] args)
    {

      Process[] processes = [
        new Process { ArrivalTime = 0, BurstTime = 1, Priority = 2},
        new Process { ArrivalTime = 1, BurstTime = 7, Priority = 3},
        new Process { ArrivalTime = 2, BurstTime = 22, Priority = 12},
        new Process { ArrivalTime = 3, BurstTime = 3, Priority = 5},
        new Process { ArrivalTime = 4, BurstTime = 9, Priority = 11},
        new Process { ArrivalTime = 5, BurstTime = 15, Priority = 9},
        new Process { ArrivalTime = 6, BurstTime = 26, Priority = 7},
        new Process { ArrivalTime = 7, BurstTime = 17, Priority = 6},
        new Process { ArrivalTime = 8, BurstTime = 8, Priority = 1},
        new Process { ArrivalTime = 9, BurstTime = 19, Priority = 8},
      ];

      SchedulingResult fcfcResult = FirstComeFirstServeAlgorithm.Execute(processes);
      SchedulingResult longestJobFirstResult = LongestJobFirst.Execute(processes);
      Console.WriteLine($"First Come First Serve Algo result: {fcfcResult}");
      Console.WriteLine($"Longest Job First Algo result: {longestJobFirstResult}");
    }
  }
}
