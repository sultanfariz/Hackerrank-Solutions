defmodule Solution do
  defp sum(a, b), do: a + b

  def main() do
    a = IO.gets("") |> String.trim |> String.to_integer
    b = IO.gets("") |> String.trim |> String.to_integer
    sum(a, b) |> IO.puts
  end
end

Solution.main()
