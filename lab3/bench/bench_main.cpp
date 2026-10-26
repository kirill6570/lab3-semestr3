#include <benchmark/benchmark.h>
#include "DynamicArray.h"
#include "SinglyList.h"
#include "Stack.h"
#include "Queue.h"
#include "HashTable.h"
#include "CompleteBinaryTree.h"

static void BM_DynamicArray_PushBack(benchmark::State& state) {
    for (auto _ : state) {
        DynamicArray arr;
        for (int i = 0; i < state.range(0); ++i) {
            arr.pushBack("x");
        }
    }
}
BENCHMARK(BM_DynamicArray_PushBack)->Range(1000, 64000);

static void BM_SinglyList_PushBack(benchmark::State& state) {
    for (auto _ : state) {
        SinglyList list;
        for (int i = 0; i < state.range(0); ++i) {
            list.pushBack("x");
        }
    }
}
BENCHMARK(BM_SinglyList_PushBack)->Range(1000, 64000);

static void BM_Stack_Push(benchmark::State& state) {
    for (auto _ : state) {
        Stack s;
        for (int i = 0; i < state.range(0); ++i) {
            s.push("x");
        }
    }
}
BENCHMARK(BM_Stack_Push)->Range(1000, 64000);

static void BM_Queue_Push(benchmark::State& state) {
    for (auto _ : state) {
        Queue q;
        for (int i = 0; i < state.range(0); ++i) {
            q.push("x");
        }
    }
}
BENCHMARK(BM_Queue_Push)->Range(1000, 64000);

BENCHMARK_MAIN();