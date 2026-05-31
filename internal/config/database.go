package config

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectToDB(connString string) (*pgxpool.Pool, error) {
	
	pool, err := pgxpool.New(context.Background(), connString) //context.Background() — это пустой, ничего не делающий "пульт". Он означает: "жди столько, сколько нужно, не отменяй, не ставь таймаут".
	if err != nil {
		fmt.Println("ConnectToDB", err)
		return nil, fmt.Errorf("ConnectToDB: %w", err)
	}
//context.Background() - пустой контекст 
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) //создаем кон
	defer cancel() //defer cancel() освобождает ресурсы контекста при выходе из функции: Канал Таймер

	err = pool.Ping(ctx)

	if err != nil {
    pool.Close()
    return nil, fmt.Errorf("база не отвечает: %w", err)
}
	
	return pool, nil

}



//Что произойдёт, если вернуть pgxpool.Pool (значение)

//Внутри функции создаётся оригинал pool (с мьютексами, соединениями, счётчиками).
//При return pool Go создаёт поверхностную копию всей структуры pgxpool.Pool. Поля, которые являются числами или указателями, копируются. Но мьютекс (sync.Mutex) — это структура с внутренним состоянием; при копировании он становится независимой копией, но при этом оба мьютекса (оригинал и копия) контролируют одни и те же сетевые соединения.
//Оригинал внутри функции уничтожается (выход из функции). Но его мьютексы уже не используются, а копия (которую ты получил снаружи) имеет свои мьютексы.
//Когда ты вызываешь метод пула (например, pool.Query(...)), он блокирует свою копию мьютекса, но соединения уже управляются двумя разными мьютексами — синхронизация ломается. Результат: гонки данных, паники, некорректная работа.