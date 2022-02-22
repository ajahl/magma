#include <sentry.h>

#include <iostream>
#include <pwd.h>
#include <string.h>
#include <signal.h>
#include <unistd.h>

void crash1() { 
    volatile int* a = (int*)(NULL); 
    *a = 1; 
}

int crash2(int i) {
        i = i / 0;
        return i;
}

static sentry_value_t before_send(sentry_value_t event, void *hint, void *data) {
  /* sentry_value_t functions handle NULL automatically */
  std::cout << "==== sentry callback ====\n";
//   sentry_value_t exceptions = sentry_value_get_by_key(event, "exception");
//   sentry_value_t values = sentry_value_get_by_key(exceptions, "values");
//   sentry_value_t exception = sentry_value_get_by_index(values, 0);
//   sentry_value_t type = sentry_value_get_by_key(exception, "type");
//   const char *type_str = sentry_value_as_string(type);
//   std::cout << type_str << std::endl;

//   /* use the data passed during initialization */
//   const char *custom_error = (const char *)hint;

//   if (type_str && strcmp(type_str, custom_error) == 0) {
//     sentry_value_t fingerprint = sentry_value_new_list();
//     sentry_value_append(fingerprint, sentry_value_new_string("custom-error"));
//     sentry_value_set_by_key(event, "fingerprint", fingerprint);
//   }
//   return event;
  return sentry_value_new_null();
}

int main() {
    sentry_options_t* options = sentry_options_new();
    // sentry_options_set_before_send(options, before_send, (void *)"CustomError");
    sentry_options_set_dsn(options, "...");
    sentry_options_set_release(options, "test@2.3.12");
    sentry_options_set_database_path(options, "sentry-db");
    sentry_options_set_system_crash_reporter_enabled(options, 1);
    sentry_options_set_symbolize_stacktraces(options, 1);
    sentry_options_add_attachment(options, "/var/log/mme.log");
    sentry_options_add_attachment(options, "/var/log/syslog");
    sentry_options_set_debug(options, 1);
    sentry_init(options);
    sentry_set_transaction("init");
     
    int i = 123;
    std::cout << "==== sentry crasher ====\n";

    while(true){  
        char ch;
        std::cout << "input:";
        std::cin >> ch;  
        std::cout << "ch = " << ch << "\n";
        std::cout << "new i = " << i << "\n";    
        // i = crash2(i);
        crash1();
    }
    sentry_shutdown();
    return(0);
}
