# Task for Avito Core-Services Internship

## Main Task

* Get list of locations
* Get tree of categories
* Count number of items in each category in Moscow

## Optional

* Get statistics for any city
* HumanReadable statistics

## Installation (create an image)

        git clone https://github.com/koliankolin/avito_core.git
        cd avito_core
        docker build -t go_app .
        
## Examples of Usage

### Locations

* #### Save all locations in JSON
        docker run --rm -v $(pwd)/data:/go/data go_app parser get locations all
        // Possible output:
        // Your JSON-file in ./data/locationsTree.json
        
* #### Save one location in JSON (works only with regions not cities)
        docker run --rm -v $(pwd)/data:/go/data go_app parser get locations -n 'Ярославская область'
        // Possible output:
        // Your JSON-file in ./data/Yaroslavskaya Oblast′

* #### Print all locations in terminal
        docker run --rm go_app parser get locations all -p
        // Possible output:
        // Россия
           	├── Тульская область
           		├── Алексин
           		├── Агеево
           		├── Арсеньево
           		├── Богородицк
           		├── Белев
           		├── Болохово
           		├── Барсуки
           		├── Брусянский
           		├── Бородинский
           		├── Венев
           		├── Волово
           		├── Грицовский
           	...

* #### Print one location in terminal (works only with regions not cities)
        docker run --rm go_app parser get locations -n 'Свердловская область' -p
        // Possible output:
        // Свердловская область
                   ├── Асбест
                   ├── Алапаевск
                   ├── Артемовский
                   ├── Арамиль
                   ├── Азанка
                   ├── Арти
                   ├── Ачит
                   ├── Атиг
                   ├── Березовский
                ...

### Categories

* #### Save all categories in JSON
        docker run --rm -v $(pwd)/data:/go/data go_app parser get categories
        // Possible output:
        // Your JSON-file in ./data/categories.json
        
* #### Print all categories in terminal
        docker run --rm go_app parser get categories -p
        // Possible output:
        // Недвижимость
                   ├── Квартиры
                   ├── Комнаты
                   ├── Дома, дачи, коттеджи
                   ├── Земельные участки
                   ├── Гаражи и машиноместа
                   └── Недвижимость за рубежом
           Бытовая электроника
                   ├── Аудио и видео
                   ├── Игры, приставки и программы
                   ├── Настольные компьютеры
                   ├── Ноутбуки
                ...

### Statistics

* #### Save total statistics in JSON
        docker run --rm -v $(pwd)/data:/go/data go_app parser get statistics all
        // Possible output:
        // Your JSON-file in /go/data/statisticsTotal.json

* #### Save one location's statistics in JSON
        docker run --rm -v $(pwd)/data:/go/data go_app parser get statistics -n москва
        // Possible output:
        // Your JSON-file in ./data/statisticsMoskva.json

* #### Print total statistics in terminal
        docker run --rm go_app parser get statistics all -p
        // Possible output:
        // Локация: Россия
           Общее количество объявлений: 50855435
           Для дома и дачи: 5354345
           Бытовая электроника: 3903271
           Хобби и отдых: 3816255
           Услуги: 1397385
           Личные вещи: 19334177
           Транспорт: 11168411
           Недвижимость: 2714718
           Работа: 2105744
           Животные: 577424
           Для бизнеса: 486886

     
* #### Print one location's statistics in terminal
        docker run --rm go_app parser get statistics -n Санкт-Петербург -p
        // Possible output:
        // Локация: Санкт-Петербург
           Общее количество объявлений: 4578503
           Работа: 129076
           Для бизнеса: 33995
           Недвижимость: 97874
           Услуги: 82841
           Животные: 30155
           Личные вещи: 1948707
           Транспорт: 993194
           Для дома и дачи: 453548
           Хобби и отдых: 447677
           Бытовая электроника: 361394
           
