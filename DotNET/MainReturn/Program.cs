using System;

namespace ReturnValue;

//// Function Main Role
// 1. Main는 `static`이 아닌 멤버는 사용할 수 없다.
// 2. 액세스 한정자를 소유할 수 있다. <Public, Private, Internal, ...etc>
// 3. `void int Task`나 `Task<T>`을 반환할 수 있다.
class Program
{
    static string ReturnValue(string s)
    {
        return s;
    }

    // `Task` 또는 `Task<T>`는 타입의 결과를 반환하는 작업
    // 비동기 메서드에서 사용되며, `async`와 `await` 키워드를 사용하여 비동기 작업이 완료될 때까지 대기 후 처리
    static async Task<string> AsyncReturnValue(string s)
    {
        // `Task.Delay(int)`는 지정된 `n` 시간동안 비동기적으로 대기하는 작업을 생성
        await Task.Delay(100);
        return s;
    }

    static async Task<string> OuputAsyncValue()
    {
        return await AsyncReturnValue("Hello World");
    }

    static void TypeChangedValue()
    {   
        // `{Datatype}.Parse`는 문자열을 해당 타입으로 변환, 예외가 발생할 수 있으므로 입력이 유효한지 확인 필요
        Console.WriteLine($"Type Changed Value of Parse: {int.Parse("123")}");
        
        // `Convert.{Datatype}`은 문자열을 해당 타입으로 변환, `Parse`와 유사하지만 예외 처리 방식이 다름
        Console.WriteLine($"Type Changed Value of Convert: {Convert.ToInt32("123")}");
    }

    static async Task Main(string[] args)
    {
        Console.WriteLine(ReturnValue("Hello World"));
        Console.WriteLine(await OuputAsyncValue());
    
        TypeChangedValue();
        Example.ExampleMain();
    }
}

public class Example
{
    public string GetString(int ID)
    {
        if (ID < names.Length)
        {
            return names[ID];
        }
        else
        {
            return "ID is out of range.";
        }
    }

    private string[] names = { "Alice", "Bob", "Charlie" };

    public static void ExampleMain()
    {
        Example example = new Example();
        Console.WriteLine(example.GetString(1)); // Output: Bob
        Console.WriteLine(example.GetString(5)); // Output: ID is out of range.
    }
}