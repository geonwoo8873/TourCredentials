# Rand libary

## Updated change logic

* `~ 0.8.0`

    ```rs
    use rand::Rng();

    let random_value = rand::thread_rng().get_renge(start..end);
    ```

* `0.9.0 ~`
  
    ```rs
    use rand::prelude::*;

    let random_value = rand::rng().random_range(start..end);
    // Support type : uninteger, integer, float<32/64>
    let random_data = rng.random::<Type>();
    ```