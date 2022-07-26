#include<iostream>
#include"redis.h"
using namespace std;

int main()
{
    MyRedis rdb;
    rdb.connect();
    rdb.setString("moyu","include<iostream>\nusing namespace std;\nint main()\n{\n    cout<<\"hello world\"<<endl;\n    return 0;\n}\n");
    cout<<rdb.getString("moyu")<<endl;
    rdb.setbit("test",123,1);
    rdb.setbit("test",124,1);
    cout<<rdb.getbit("test",123)<<endl;
    cout<<rdb.getbit("test",1)<<endl;
    cout<<rdb.bitcount("test")<<endl;
    rdb.lpush("he","123");
    rdb.lpush("he","1234");
    string res;
    do
    {
        res = rdb.rpop("he");
        if(res.size())  cout<<res<<endl;
    } while (!res.empty());
    rdb.close();
    return 0;
}
