# go-course-for-students

Ваши домашние задания будут проверяться только при условии, если они сдаются в виде pr-а внутри приватного форка.

### Как сделать приватный форк

1. Создайте голый клон репозитория
    ```bash
    git clone --bare git@github.com:Tinkoff/go-course-for-students.git
    ```

2. [Создайте новый приватный репозиторий на github](https://docs.github.com/ru/repositories/creating-and-managing-repositories/creating-a-new-repository)
   и назовите его go-course-for-students

3. Зеркально отправьте свой голый клон в новый go-course-for-students репозиторий
   > Замените `<your_username>` своим именем на Github в приведенном ниже URL-адресе

    ```bash
    cd go-course-for-students.git
    git push --mirror git@github.com:<your_username>/go-course-for-students.git
    ```

4. Удалите временный локальный репозиторий, созданный на шаге 1
   ```bash
   cd ..
   rm -rf go-course-for-student.git
   ```

5. Теперь вы можете клонировать репозиторий `go-course-for-students` на свой компьютер
   ```bash
   git clone git@github.com:<your_username>/go-course-for-students.git
   ```

6. Для того чтобы мы могли проверить ДЗ надо [предоставить доступ в репозиторий](https://docs.github.com/ru/account-and-profile/setting-up-and-managing-your-personal-account-on-github/managing-access-to-your-personal-repositories/inviting-collaborators-to-a-personal-repository)
пользователю [GoCourseTeachers](https://github.com/GoCourseTeachers)


### Как подтянуть изменения в форк

Обратите внимание, для того чтобы скачать спецификацию и тесты локально нужно подтянуть изменения из основного репозитория.
Для этого:
1. Замержите в свой main всю накопленную работу в своём форке и переключитесь на обновлённый main локально
2. Если не настроен upstream, то сделайте ```git remote add upstream git@github.com:Tinkoff/go-course-for-students.git``` или ```git remote add upstream https://github.com/Tinkoff/go-course-for-students.git```
3. Обновите upstream: ```git fetch upstream```
4. Подтяните изменения из upstream и ребазируйтесь на них: ```git rebase upstream/main```
