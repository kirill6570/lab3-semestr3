#ifndef TEXTSERIALIZER_H
#define TEXTSERIALIZER_H

#include <fstream>
#include <string>
#include <iostream>
#include "SinglyList.h"

class TextSerializer {
public:
    static void saveText(const SinglyList& list, const std::string& filename) {
        std::ofstream out(filename);
        if (!out) {
            std::cerr << "Не удалось открыть файл для записи" << std::endl;
            return;
        }

        out << list.getSize() << "\n";
        for (int i = 0; i < list.getSize(); ++i) {
            out << list.get(i) << "\n";
        }

        out.close();
        std::cout << "Сериализовано в " << filename << " (текстовый)" << std::endl;
    }

    static SinglyList loadText(const std::string& filename) {
        std::ifstream in(filename);
        if (!in) {
            std::cerr << "Не удалось открыть файл для чтения" << std::endl;
            return SinglyList();
        }

        int size = 0;
        in >> size;
        in.ignore();

        SinglyList list;
        for (int i = 0; i < size; ++i) {
            std::string line;
            std::getline(in, line);
            list.pushBack(line);
        }

        in.close();
        std::cout << "Десериализовано из " << filename << " (текстовый)" << std::endl;
        return list;
    }
};

#endif