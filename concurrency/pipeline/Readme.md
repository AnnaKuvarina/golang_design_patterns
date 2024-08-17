## Pipeline pattern
Involves breaking down the task into the stages, where each stage is a goroutine that process data and passes it to the next  stage though the channel.
This pattern allows for efficient and scalable concurrent processing, as stages can operate independent and simultaneously.
