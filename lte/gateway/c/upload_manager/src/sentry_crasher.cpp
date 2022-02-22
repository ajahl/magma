#include <sentry.h>
// #include <syslog.h>
#include <iostream>
#include <pwd.h>
#include <string.h>
#include <signal.h>
#include <unistd.h>

#define SYS_LOG_PATH "/tmp/magma_services.log"

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
//   char *command = (char*) "journalctl -u magma@* -n 1000 > /tmp/magma.log";
//   system(command);
//   return event;
  std::string command = "journalctl -u magma@* -n 1000 > ";
  command.append(SYS_LOG_PATH);
  // char *command = (char*) "journalctl -u magma@* -n 1000 > /tmp/magma.log";
  system(command.c_str());
  return event;
}

int main() {
    sentry_options_t* options = sentry_options_new();
    sentry_options_set_before_send(options, before_send, (void *)"CustomError");
    sentry_options_set_dsn(options, "https://cb8f04b2aa1a4ce1ae42e9a6a5b88b34@o1036448.ingest.sentry.io/6169376");
    sentry_options_set_release(options, "test@2.3.12");
    sentry_options_set_database_path(options, "sentry-db");
    sentry_options_set_system_crash_reporter_enabled(options, 1);
    sentry_options_set_symbolize_stacktraces(options, 1);
    sentry_options_add_attachment(options, "/var/log/mme.log");
    sentry_options_add_attachment(options, SYS_LOG_PATH);
    // sentry_options_set_debug(options, 1);
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
