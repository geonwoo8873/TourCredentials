using System;

class Program
{
    private const string Line = "------------------------------";

    string ReturnValue(string s)
    {
        return s;
    }
    
    // `ToUpper`는 문자열을 대문자로 변환하는 메서드
    // `ToLower`는 문자열을 소문자로 변환하는 메서드
    static void UpperFormat()
    {
        Program program = new Program();
        string result = program.ReturnValue("This is a string.");

        Console.WriteLine($"""
        {Line}
        Before : {result}
        After Upper : {$"[{result.ToUpper()}]"}
        After Lower : {$"[{result.ToLower()}]"}
        """);
    }

    // `Trim`은 문자열의 시작과 끝에서 공백을 제거
    // `TrimStart`는 문자열의 시작에서 공백을 제거
    // `TrimEnd`는 문자열의 끝에서 공백을 제거
    static void TrimFormat()
    {
        Program program = new Program();
        string result = program.ReturnValue("      The Trim method string.     ");

        Console.WriteLine($"""
        {Line}
        Before : {result}
        After Normal Trim : {$"[{result.Trim()}]"}
        After Trim Start  : {$"[{result.TrimStart()}]"}
        After Trim End    : {$"[{result.TrimEnd()}]"}
        """);
    }
    
    // `Replace`는 해당 문자열에서 특정 부분을 찾고 다른 문자열 값으로 변경 <Replace(Value, ReplaceValue)>
    static void ReplaceFormat()
    {
        Program program = new Program();
        string result = program.ReturnValue("The Replace method string.");

        Console.WriteLine($"""
        {Line}
        Before : {result}
        After Replace : {$"[{result.Replace("Replace", "Replaced")}]"}
        """);
    }

    // `Contains`는 문자열 값에 특정 문자열이 포함되어 있으면 `true`를 반환
    static void SearchFormat()
    {
        Program program = new Program();
        string result = program.ReturnValue("The Search method string.");
        Console.WriteLine($"""
        {Line}
        Before : {result}
        Get Search : {$"[{result.Contains("Search")}]"}
        Get Search : {$"[{result.Contains("search")}]"}
        """);
    }

    // `StartsWith`는 문자열이 특정 문자열로 시작하는지 여부를 반환
    // `EndsWith`는 문자열이 특정 문자열로 끝나는지 여부를 반환
    static void WithFormat()
    {
        Program program = new Program();
        string result = program.ReturnValue("The With method string.");
        
        Console.WriteLine($"""
        {Line}
        Before : {result}
        Starts With : {$"[{result.StartsWith("With")}]"}
        Ends With : {$"[{result.EndsWith("string.")}]"}
        """);
    }

    static void Main()
    {
        UpperFormat();
        TrimFormat();
        ReplaceFormat();
        SearchFormat();
        WithFormat();
    }
}