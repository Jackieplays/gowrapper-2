#include "test1.h"
#include <stdio.h>

void prime(int n)
{
    int i, flag = 0;
 
    printf("Enter a positive integer:\n ");
    
 
    for(i = 2; i <= n/2; ++i)
    {
        
        if(n%i == 0)
        {
            flag = 1;
            break;
        }
    }
 
    if (n == 1) 
    {
      printf("1 is neither a prime nor a composite number.\n");
    }
    else 
    {
        if (flag == 0)
          printf("%d is a prime number.\n", n);
        else
          printf("%d is not a prime number.\n", n);
    }
 
}
