# Biocad_test
Необходимо было написать сервис на языке golang который будет реализовывать следующие функции: 
* брать конфиги(host,port,password и т.д.) для соединения с БД (бд на выбор) а так же адрес директории
* периодически осматривать директорию на наличие новых не обработанных еще файлов (.tsv) (вероятно держать в базе список уже обработанных) (см. лист "Исходные данные")
* обработку файлов ставить в очередь 
* обработка файла - нужно распарсить .tsv и положить в соответствующую структуру (формат файла статичный, поля/количество не меняется)
* данные из файла поместить в БД
* после обработки файла и записи в БД нужно сформировать файл(rtf,doc,pdf на выбор) с названием из поля *unit_guid* в входном файле, с данными по этому *unit_guid*
* ошибки парсинга (например не соответсвие файла) - тоже записывать в БД и файл
* выходные файлы размещать в отдельной директории
* сделать API-интерфейс который позволит получать из БД данные с пагинацией (page/limit) для получения данных по *unit_guid*

Примечание: *unit_guid* - это некий уникальный идентификатор устройства, в файле содержатся строки описывающие множество устройств

