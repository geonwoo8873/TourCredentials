using System;
using Npgsql;

namespace PostgreSQL;

class Program
{
    async static Task Main()
    {
        var connString = Environment.GetEnvironmentVariable("PG_CONN_STRING")
            ?? "Host=localhost;Port=9999;Username=postgres;Password=9563;Database=postgres";

        var dataSourceBuilder = new NpgsqlDataSourceBuilder(connString);
        await using var dataSource = dataSourceBuilder.Build();
        await using var conn = await dataSource.OpenConnectionAsync();

        // Ensure demo table exists before insert/select.
        await using (var createCmd = new NpgsqlCommand(
            """
            CREATE TABLE IF NOT EXISTS data (
                id SERIAL PRIMARY KEY,
                some_field TEXT NOT NULL
            );
            """, conn))
        {
            await createCmd.ExecuteNonQueryAsync();
        }

        await using (var insertCmd = new NpgsqlCommand("INSERT INTO data (some_field) VALUES (@p)", conn))
        {
            insertCmd.Parameters.AddWithValue("p", "Hello world");
            await insertCmd.ExecuteNonQueryAsync();
        }

        await using (var selectCmd = new NpgsqlCommand("SELECT some_field FROM data", conn))
        await using (var reader = await selectCmd.ExecuteReaderAsync())
        {
            while (await reader.ReadAsync())
            {
                Console.WriteLine(reader.GetString(0));
            }
        }
    }
}