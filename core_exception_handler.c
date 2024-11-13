
extern void gocv_exception_handler(char*,char*);

int cb_gocv_exception_handler(int status, const char *func_name, const char *err_msg, const char *file_name, int line, void *userdata) {
	gocv_exception_handler((char*)func_name ,(char*)err_msg);
}

