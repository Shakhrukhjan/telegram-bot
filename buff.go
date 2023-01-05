#include <iostream>
#include <vector>
#include <fstream>
#include <sstream>
using namespace std;
struct LST {
    int val;
    int index;
    LST* next;
};
template <class Type>
class List {
    LST* first;
    public:
        List() {
            first = NULL;
        }
void add(int x) {    //Функция для добавления элемента в список
        LST* temp = new LST;
        temp->val = x;
        temp->next = NULL;
        if (first == NULL) {
            first = temp;
        } else {
            LST* tempNew = first;
            while (tempNew-> next != NULL) {
                tempNew = tempNew->next;
            }
            tempNew->next = temp;
        }
}
void insert(int x, int pos) { //Функция добавления элемента в заданную позицию
        if (pos == 1) {
            LST* temp = new LST;
            temp->val = x;
            temp->next = first;
            first = temp;
        } else {
            LST* temp = new LST;
            temp->val = x;
            if (first == NULL) {
                first = temp;
            } else {
                LST* tempNew = first;
                int count = 1;
                while (tempNew->next != NULL && count < pos - 1) {
                    tempNew = tempNew->next;
                    count++;
                }
                LST* tempNew2 = tempNew->next;
                tempNew->next = temp;
                temp->next = tempNew2;
            }
        }
}
void delByPos(int pos) { //Функция удаления элемента под заданной позицией
        LST* temp = first;
        if (pos == 1) {
            first = first->next;
            delete temp;
        } else {
            LST* del;
            int count = 1;
            while (first != NULL && count < pos - 1) {
                temp = temp->next;
                count++;
            }
            del = temp->next;
            temp->next = del->next;
            delete del;
        }
}
void delByVal(int value) { //Функция удаления элемента по его значению
        LST* temp = first;
        LST* del = first->next;
        if (temp->val == value) {
            first = first->next;
            delete temp;
        } else {
            while (del != NULL) {
                if (del->val == value) {
                    break;
                } else {
                    temp = del;
                    del = del->next;
                }
            }
            if (del == NULL) {} else {
                temp->next = del->next;
                delete del;
            }
        }
}
void print() {         //Функция вывода списка
        LST* temp = first;
        while (temp != NULL) {
            cout << temp->index << "      ";
            temp = temp->next ;
        }
}
int countLST() {   //Функция определения количества элементов списка
        LST* temp = first;
        int countL = 0;
      while(temp != NULL) {
          countL++;
          temp = temp->next;
       }
    return countL;
}
string lstToString(){
  stringstream os;
  LST* temp = first;
        while (temp != NULL) {
            os << temp->val << " ";
            temp = temp->next;
        }
 return os.str();
}

void saveToFile(){     // сохранить элементы списка в файла
   ofstream fout("listOne.txt");
    LST* temp = first;
        while (temp != NULL) {
            fout << temp->val << " ";
            temp = temp->next;
        }
        fout.close();
}
void loadToFile() {      // вывод элементы списка из файла

    ifstream fin("listOne.txt");
     if (fin) {
        int number = 0;
        while(fin >> number) {
            add(number);
        }
        cout << "Элементы списка"<<endl;
     }
     else {
        cout << "ERROR" ;
     }
}

void reIndex(){
    LST *buff = first;
    int k=0;
    while(buff->next != NULL) {
        buff->index = k;
        k++;
        buff = buff -> next;
    }
    buff -> index = k;
}

Type& operator[](int index) {
      LST* temp = first;

        while (temp->next != NULL) {
            if (index==temp->index) return temp->val;
            temp = temp->next;
        }
        if (index==temp->index) return temp->val;
}
void sortLST(List<Type>t){               // сортировка элементы списка
       t.reIndex();
       for (int i=0; i<countLST(); i++) {
         for (int j=i; j<countLST(); j++) {
            if (t[i] < t[j]) {
                swap(t[i],t[j]);
            }
         }
       }
       for (int i=0; i<countLST(); i++) {
        cout << t[i] << " ";
       }
}

};

int main() {
    float n,a;
    setlocale(LC_ALL,"rus");
    List<int> lst;
//   cout << "n = ";
//   cin >> n;
//   for (int i=0; i<n; i++) {
//    cin >> a;
//    lst.add(a);
//   }
//   cout <<"------------"<<endl;
        //lst.sortLST();
        //lst.delByPos(5);
        //int x; cout<<"element: ";  cin >> x;  lst.insert(x, 4);
        //lst.delByVal(4);
        //lst.reIndex();
        //int curr = lst.countLST();
        //cout<<lst.countLST()<<endl;
        lst.loadToFile();  lst.print();
         //lst.print();
        //cout<<lst.lstToString();
        //lst.saveToFile();
       // lst.sortLST(lst);
        // cout << lst[3];
        return 0;
}


--------------------------------------------
#include <iostream>
#include <vector>
#include <string.h>
using namespace std;
struct Tourist {
   string name_tour;             // наименование тура
   string fam_client;            // фамилия клиента
   float price_day;              // цена за день
   float amount_day;             // количество дней
   float fare;                   // стоимость проезда
   string exch_curr;             // курс валюта
   float amount_curr;            // количество валюты
   float total_price;            // стоимость проездка
};
void findd(vector<Tourist>tour) {        // Поиск по наименование тура
     int n,p=1;
     while (n!=0){
            cout<<"Поиск по:"<<endl;
      cout<<"1  Наименование тура"<<endl;
      cout<<"2  Фамилия клиента"<<endl;
      cout<<"3  цена за день"<<endl;
      cout<<"4  количество дней"<<endl;
      cout<<"5  стоимость проезда"<<endl;
      cout<<"6  курс валюта"<<endl;
      cout<<"7  количество валюты"<<endl;
      cout<<"8  стоимость проездка"<<endl;
      cout<<"0  Выход"<<endl;
      cout<<"----------------------"<<endl;
      cin>>n;
     switch(n){
         case 0: {return;}break;

  case 1:{
          string name_tour;
      int num=0;
      cout << "Введите наименование тура: ";
        cin>>name_tour;
     // getline(cin,name_tour);
    for (int i=0; i<tour.size(); i++) {
            if (name_tour == tour[i].name_tour) {
                num++;
       cout <<"наименование тура:  "<<tour[i].name_tour<<endl;
       cout<<"фамилия клиента:  "<<tour[i].fam_client<<endl;
       cout<<"цена за день:  "<<tour[i].price_day<<endl;
       cout<<"количество дней:  "<<tour[i].amount_day<<endl;
       cout<<"стоимость проезда:  "<<tour[i].fare<<endl;
       cout<<"курс валюта:  "<<tour[i].exch_curr<<endl;
       cout<<"количество валюты:  "<<tour[i].amount_curr<<endl;
       cout<<"стоимость проездка:  "<<tour[i].total_price<<endl;
       cout << endl;
        }
}
   if (num == 0) {
             cout << "Не найдено: "<<endl;
        }
         } break;
 case 2: {
          string fam_client;
      int num=0;
      cout << "Введите фамилия клиента: ";
      cin >> fam_client;
      cout<<"================"<<endl;
    for (int i=0; i<tour.size(); i++) {
            if (fam_client == tour[i].fam_client) {
                num++;
       cout <<"наименование тура:  "<<tour[i].name_tour<<endl;
       cout<<"фамилия клиента:  "<<tour[i].fam_client<<endl;
       cout<<"цена за день:  "<<tour[i].price_day<<endl;
       cout<<"количество дней:  "<<tour[i].amount_day<<endl;
       cout<<"стоимость проезда:  "<<tour[i].fare<<endl;
       cout<<"курс валюта:  "<<tour[i].exch_curr<<endl;
       cout<<"количество валюты:  "<<tour[i].amount_curr<<endl;
       cout<<"стоимость проездка:  "<<tour[i].total_price<<endl;
       cout << endl;
        }
}
   if (num == 0) {
             cout << "Не найдено: "<<endl;
        }
         } break;
 case 3:{
          float price_day;
      int num=0;
      cout << "цена за день:  ";
      cin >> price_day;
    for (int i=0; i<tour.size(); i++) {
            if (price_day == tour[i].price_day) {
                num++;
       cout <<"наименование тура:  "<<tour[i].name_tour<<endl;
       cout<<"фамилия клиента:  "<<tour[i].fam_client<<endl;
       cout<<"цена за день:  "<<tour[i].price_day<<endl;
       cout<<"количество дней:  "<<tour[i].amount_day<<endl;
       cout<<"стоимость проезда:  "<<tour[i].fare<<endl;
       cout<<"курс валюта:  "<<tour[i].exch_curr<<endl;
       cout<<"количество валюты:  "<<tour[i].amount_curr<<endl;
       cout<<"стоимость проездка:  "<<tour[i].total_price<<endl;
       cout << endl;
        }
}
   if (num == 0) {
             cout << "Не найдено: "<<endl;
        }
         } break;
 case 4:{
          float amount_day;
      int num=0;
      cout << "количество дней:  ";
      cin >> amount_day;
    for (int i=0; i<tour.size(); i++) {
            if (amount_day == tour[i].amount_day) {
                num++;
       cout <<"наименование тура:  "<<tour[i].name_tour<<endl;
       cout<<"фамилия клиента:  "<<tour[i].fam_client<<endl;
       cout<<"цена за день:  "<<tour[i].price_day<<endl;
       cout<<"количество дней:  "<<tour[i].amount_day<<endl;
       cout<<"стоимость проезда:  "<<tour[i].fare<<endl;
       cout<<"курс валюта:  "<<tour[i].exch_curr<<endl;
       cout<<"количество валюты:  "<<tour[i].amount_curr<<endl;
       cout<<"стоимость проездка:  "<<tour[i].total_price<<endl;
       cout << endl;
        }
}
   if (num == 0) {
             cout << "Не найдено: "<<endl;
        }
         } break;
 case 5:{
          float fare;
      int num=0;
      cout << "Стоимость проезда:  ";
      cin >> fare;
    for (int i=0; i<tour.size(); i++) {
            if (fare == tour[i].fare) {
                num++;
       cout <<"наименование тура:  "<<tour[i].name_tour<<endl;
       cout<<"фамилия клиента:  "<<tour[i].fam_client<<endl;
       cout<<"цена за день:  "<<tour[i].price_day<<endl;
       cout<<"количество дней:  "<<tour[i].amount_day<<endl;
       cout<<"стоимость проезда:  "<<tour[i].fare<<endl;
       cout<<"курс валюта:  "<<tour[i].exch_curr<<endl;
       cout<<"количество валюты:  "<<tour[i].amount_curr<<endl;
       cout<<"стоимость проездка:  "<<tour[i].total_price<<endl;
       cout << endl;
        }
}
   if (num == 0) {
             cout << "Не найдено: "<<endl;
        }
         } break;
 case 6:{
          string exch_curr;
      int num=0;
      cout << "курс валюта:  ";
      cin >> exch_curr;
    for (int i=0; i<tour.size(); i++) {
            if (exch_curr == tour[i].exch_curr) {
                num++;
       cout <<"наименование тура:  "<<tour[i].name_tour<<endl;
       cout<<"фамилия клиента:  "<<tour[i].fam_client<<endl;
       cout<<"цена за день:  "<<tour[i].price_day<<endl;
       cout<<"количество дней:  "<<tour[i].amount_day<<endl;
       cout<<"стоимость проезда:  "<<tour[i].fare<<endl;
       cout<<"курс валюта:  "<<tour[i].exch_curr<<endl;
       cout<<"количество валюты:  "<<tour[i].amount_curr<<endl;
       cout<<"стоимость проездка:  "<<tour[i].total_price<<endl;
       cout << endl;
        }
}
   if (num == 0) {
             cout << "Не найдено: "<<endl;
        }
         } break;
 case 7:{
         float amount_curr;
      int num=0;
      cout << "количество валюты:  ";
      cin >> amount_curr;
    for (int i=0; i<tour.size(); i++) {
            if (amount_curr == tour[i].amount_curr) {
                num++;
       cout <<"наименование тура:  "<<tour[i].name_tour<<endl;
       cout<<"фамилия клиента:  "<<tour[i].fam_client<<endl;
       cout<<"цена за день:  "<<tour[i].price_day<<endl;
       cout<<"количество дней:  "<<tour[i].amount_day<<endl;
       cout<<"стоимость проезда:  "<<tour[i].fare<<endl;
       cout<<"курс валюта:  "<<tour[i].exch_curr<<endl;
       cout<<"количество валюты:  "<<tour[i].amount_curr<<endl;
       cout<<"стоимость проездка:  "<<tour[i].total_price<<endl;
       cout << endl;
        }
}
   if (num == 0) {
             cout << "Не найдено: "<<endl;
        }
         } break;
 case 8:{
        float total_price;
      int num=0;
      cout << "стоимость проездка:  ";
      cin >> total_price;
    for (int i=0; i<tour.size(); i++) {
            if (total_price == tour[i].total_price) {
                num++;
       cout <<"наименование тура:  "<<tour[i].name_tour<<endl;
       cout<<"фамилия клиента:  "<<tour[i].fam_client<<endl;
       cout<<"цена за день:  "<<tour[i].price_day<<endl;
       cout<<"количество дней:  "<<tour[i].amount_day<<endl;
       cout<<"стоимость проезда:  "<<tour[i].fare<<endl;
       cout<<"курс валюта:  "<<tour[i].exch_curr<<endl;
       cout<<"количество валюты:  "<<tour[i].amount_curr<<endl;
       cout<<"стоимость проездка:  "<<tour[i].total_price<<endl;
       cout << endl;
        }
}
   if (num == 0) {
             cout << "Не найдено: "<<endl;
        }
         } break;


}
     }
}
int main () {
    setlocale(LC_ALL,"rus");
    vector<Tourist>tour;
    tour.push_back({"Delphin","Abdulloh",150,5,200,"TJS",3,2000});
    tour.push_back({"Parking","Muhammdi",151,6,201,"USD",4,2001});
    tour.push_back({"MiklukhoMakalay","Ikrom",152,7,202,"RUB",5,2000});
       for (int i=0; i<tour.size(); i++) {
       cout <<"наименование тура:  "<<tour[i].name_tour<<endl;
       cout<<"фамилия клиента:  "<<tour[i].fam_client<<endl;
       cout<<"цена за день:  "<<tour[i].price_day<<endl;
       cout<<"количество дней:  "<<tour[i].amount_day<<endl;
       cout<<"стоимость проезда:  "<<tour[i].fare<<endl;
       cout<<"курс валюта:  "<<tour[i].exch_curr<<endl;
       cout<<"количество валюты:  "<<tour[i].amount_curr<<endl;
       cout<<"стоимость проездка:  "<<tour[i].total_price<<endl;
       cout << endl;
        }
    findd(tour);          // Поиск по фамилия клиента
}
---------------------------------------- 

#include <bits/stdc++.h>
#include <iostream>
#include <math.h>
using namespace std;
struct Trapesoid {
  float x1;
  float y1;
  float x2;
  float y2;
  float x3;
  float y3;
  float x4;
  float y4;
};
void valid_Trapesoid(float x1,float y1,float x2,float y2,float x3,float y3,float x4,float y4){
    if (abs(x1-x2)*abs(y1-y2) == abs(x3-x4)*abs(y3-y4)){
            cout << "фигура являеться равнобедренная трапеция"<<endl;
        }
        else cout << "Фигура не являеться равнобедренная трапеци"<<endl;
 }
 void calcul_Trapesoid(float x1,float y1,float x2,float y2,float x3,float y3,float x4,float y4) {
    float AB,BC,CD,AD,S,P;
    AB = sqrt(pow((x1-x2),2)+pow((y1-y2),2));
    CD = sqrt(pow((x3-x4),2)+pow((y3-y4),2));
    BC = sqrt(pow((x2-x3),2)+pow((y2-y3),2));
    AD = sqrt(pow((x1-x4),2)+pow((y1-y4),2));
    S = ((BC + AD)/2) * sqrt(AB*AB - ((AD - BC) * (AD - BC))/4);
    P = AB + BC + CD +AD;
 cout<< "Длина AB = "<<AB<<endl;
 cout<< "Длина CD = "<<CD<<endl;
 cout<< "Длина BC = "<<BC<<endl;
 cout<< "Длина AD = "<<AD<<endl;
 cout<< "Площад трапеция S = "<<S<<endl;
 cout<< "Периметр трапеция P = "<<P<<endl;
 }
 float sred_Trapesoid(vector<Trapesoid>t){
              float AB, BC, CD, AD, S, P, sumS = 0;
    for (int i=0; i<t.size(); i++) {
         AB = sqrt(pow((t[i].x1-t[i].x2),2)+pow((t[i].y1-t[i].y2),2));
         CD = sqrt(pow((t[i].x3-t[i].x4),2)+pow((t[i].y3-t[i].y4),2));
         BC = sqrt(pow((t[i].x2-t[i].x3),2)+pow((t[i].y2-t[i].y3),2));
         AD = sqrt(pow((t[i].x1-t[i].x4),2)+pow((t[i].y1-t[i].y4),2));
         S = ((BC + AD)/2) * sqrt(AB*AB - ((AD - BC) * (AD - BC))/4);
         cout<<"AB = "<<AB<<endl;
         cout<<"BC = "<<BC<<endl;
         cout<<"CD = "<<CD<<endl;
         cout<<"AD = "<<AD<<endl;
         cout <<"S = "<<S<<endl;
         cout<<"----------"<<endl;
        sumS = sumS + S;
    }
    return sumS/t.size();
 }
 int count_Trapesoid(vector<Trapesoid>t, float sred) {
     float AB, BC, CD, AD, S, P, sumS = 0,coun = 0;
    for (int i=0; i<t.size(); i++) {
         AB = sqrt(pow((t[i].x1-t[i].x2),2)+pow((t[i].y1-t[i].y2),2));
         CD = sqrt(pow((t[i].x3-t[i].x4),2)+pow((t[i].y3-t[i].y4),2));
         BC = sqrt(pow((t[i].x2-t[i].x3),2)+pow((t[i].y2-t[i].y3),2));
         AD = sqrt(pow((t[i].x1-t[i].x4),2)+pow((t[i].y1-t[i].y4),2));
         S = ((BC + AD)/2) * sqrt(AB*AB - ((AD - BC) * (AD - BC))/4);
         if (S > sred) {
            coun++;
         }
    }
    return coun;
 }

int main () {
      setlocale(LC_ALL,"rus");
     vector<Trapesoid>trap;
     trap.push_back({0,0,3,4,6,4,9,0});
     trap.push_back({1,0,2.5,2,5,2,6,0});
     trap.push_back({2,0,4,2.5,6,2.5,8,0});
          trap.push_back({2,1,3,5,6,5,8,1});

     calcul_Trapesoid(0,0,3,4,6,4,9,0);
     cout<<"---------------------------"<<endl;
     valid_Trapesoid(2,1,3,5,6,5,8,1);
     cout<<"---------------------------"<<endl;
     float  sred = sred_Trapesoid(trap);
     cout << count_Trapesoid(trap, sred)<<endl;
}
---------------------------------------------------
#include <iostream>
using namespace std;
struct Equation {
    double first;
    double second;
double func(double x) {   // вычисление значение y

    return first * x + second;
}
double root() {
   double koren_ur;
   koren_ur = - second / first;
   if (first == 0) {
    cout << "на 0 делить нельзя"<<endl;
   }
   else {
      return   koren_ur;
   }
}
};
int main ()
{
    setlocale(LC_ALL,"rus");
    double A,B,x;
    Equation buff;
    cout << "              Операция с уравненим вида Ax+B"<<endl;
    cout << "1  Нахож знач у при определенном коеф A, B и заданном значение x"<<endl;
    cout << "2  Нахождение корен урав"<<endl;
    cout<<"---------------------------"<<endl;
    int n; cin>>n;
         switch (n) {
    case 1:
        {
               cout<<"значение A = ";cin>>A; buff.first = A;
              cout<<"значение x = ";cin>>x;
              cout<<"значение B = ";cin>>B; buff.second = B;
              cout <<endl;
              cout<< "Result: у = " << buff.func(x) <<endl;
        }break;
    case 2:
        {
             cout<<"значение A = ";cin>>A; buff.first = A;
             cout<<"значение B = ";cin>>B; buff.second = B;
             cout <<"Result:  x = "<<buff.root() << " ";
        }break;
   }
}
