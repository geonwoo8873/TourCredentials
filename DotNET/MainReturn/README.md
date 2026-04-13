# Main Function

Main 함수에는 `static` 키워드가 함께 사용되어야 하며 다른 멤버의 키워드는 사용할 수 없다. 액세스 한정자를 소유할 수 있으며 [Public, Private, Internal, ...etc]에 포함될 수 있고 `void <Datatype> Task`나 `Task<T>`을 반환할 수 있다.

## Access group

```cs
// Default private member
static void Main() 
{
    //...
}

public static void Main()
{
    //...
}
```

## Async Task

```cs
// Async process task action
static async Task Main() 
{
    await ...
}

static async Task<T> Main() 
{
    await ...
}
```