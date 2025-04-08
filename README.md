# visibleK8S的后端项目，自用
基于同作者仓库JWT-TEST二次开发

*使用了gin jwt/v5 logrus*

main.go中本地调试时取消了跨域限制，请不要在生产环境使用

--embed-certs=true 内嵌证书

corev1 "k8s.io/api/core/v1"
metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"