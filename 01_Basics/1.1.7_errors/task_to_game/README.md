# Игра без права на ошибку
В совсем маленьких программах обертывать ошибки обычно не требуется. Поэтому сразу скажу, что задачка эта непривычно большая. Заодно попрактикуетесь в работе с унаследованным кодом на Go. Поскольку это финальная задача модуля — можно считать ее толстым боссом ツ

Мы будем писать игру, в которой любая ошибка приводит к трагическому финалу. Сначала пройдем по коду, а затем я расскажу, что требуется сделать.

### Задание
Как видно по коду, ошибки создаются через `errors.New()` и `fmt.Errorf()`. 
Хочется больше структуры, поэтому добавьте отдельный тип для каждого вида ошибок:

```go

// invalidStepError - ошибка, которая возникает,
// когда команда шага не совместима с объектом
type invalidStepError any

// notEnoughObjectsError - ошибка, которая возникает,
// когда в игре закончились объекты определенного типа
type notEnoughObjectsError

// commandLimitExceededError - ошибка, которая возникает,
// когда игрок превысил лимит на выполнение команды
type commandLimitExceededError

// objectLimitExceededError - ошибка, которая возникает,
// когда игрок превысил лимит на количество объектов
// определенного типа в инвентаре
type objectLimitExceededError
```

Метод `player.do()` должен возвращать либо одну из этих ошибок, либо `nil` (если ошибок не было).

Кроме того, создайте ошибку верхнего уровня:

```go
// gameOverError - ошибка, которая произошла в игре
type gameOverError struct {
// количество шагов, успешно выполненных
// до того, как произошла ошибка
nSteps int
// ...
}
```

Метод `game.execute()` должен возвращать либо ошибку типа `gameOverError`, либо `nil` (если ошибок не было).

* Если метод получил ошибку от `player.do()` — пусть обернет ее в gameOverError.
* Если ошибка произошла в самом `game.execute()` — пусть создаст ошибку подходящего типа, а затем обернет ее в `gameOverError`.

И последнее. Создайте функцию `giveAdvice()`, которая дает игроку совет, как избежать случившейся ошибки в будущем:

```go
func giveAdvice(err error) string {
// ...
}
```

Правила работы giveAdvice():

* Если команда не совместима с объектом, возвращает things like 'COMMAND OBJECT' are impossible, где COMMAND — название команды, а OBJECT — название объекта.
* Если в игре закончились объекты определенного типа, возвращает be careful with scarce OBJECTs, где OBJECT — название объекта.
* Если игрок слишком много ел, возвращает eat less. Если игрок слишком много говорил, возвращает talk to less.
* Если игрок превысил лимит на количество объектов определенного типа в инвентаре, возвращает don't be greedy, LIMIT OBJECT is enough, где LIMIT — значение лимита, а OBJECT — название объекта.

Например:

```
things like 'eat bob' are impossible
be careful with scarce mirrors
don't be greedy, 1 apple is enough
```

Итого:

* Создайте отдельные типы ошибок и возвращайте ошибки этих типов в подходящих случаях.
* Создайте тип gameOverError и возвращайте ошибку этого типа из game.execute().
* Создайте функцию giveAdvice(), которая возвращает совет на основе ошибки.
