#include "lte/gateway/c/core/oai/test/spgw_task/spgw_test_util.h"


class SPGWIPv4NetworkTest : public ::testing::Test {
  
  virtual void SetUp() {
    // initialize configs
    std::cout << "Setup done" << std::endl;
  }

  virtual void TearDown() {
  }

 protected:
};

TEST_F(SPGWIPv4NetworkTest, TestExtractIPv4NetworkSuccess) {
    
}