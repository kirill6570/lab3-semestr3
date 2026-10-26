#ifndef SERIALIZER_H
#define SERIALIZER_H

#include <fstream>
#include <string>
#include <iostream>
#include "DynamicArray.h"

class Serializer {
public:
    static void saveBinary(const DynamicArray& arr, const std::string& filename) {
        std::ofstream out(filename, std::ios::binary);
        if (!out) {
            std::cerr << "Не удалось открыть файл для записи" << std::endl;
            return;
        }

        int size = arr.getSize();
        out.write(reinterpret_cast<const char*>(&size), sizeof(size));

        for (int i = 0; i < size; ++i) {
            std::string val = arr.get(i);
            int len = val.size();
            out.write(reinterpret_cast<const char*>(&len), sizeof(len));
            out.write(val.data(), len);
        }

        out.close();
        std::cout << "Сериализовано в " << filename << " (бинарный)" << std::endl;
    }

    static DynamicArray loadBinary(const std::string& filename) {
        std::ifstream in(filename, std::ios::binary);
        if (!in) {
            std::cerr << "Не удалось открыть файл для чтения" << std::endl;
            return DynamicArray();
        }

        int size = 0;
        in.read(reinterpret_cast<char*>(&size), sizeof(size));

        DynamicArray arr;
        for (int i = 0; i < size; ++i) {
            int len = 0;
            in.read(reinterpret_cast<char*>(&len), sizeof(len));
            std::string val(len, '\0');
            in.read(&val[0], len);
            arr.pushBack(val);
        }

        in.close();
        std::cout << "Десериализовано из " << filename << " (бинарный)" << std::endl;
        return arr;
    }
};

#endif