# 使用官方的 NGINX 镜像
FROM nginx:alpine

# 复制自定义的 NGINX 配置文件到容器中
COPY nginx.conf /etc/nginx/conf.d/default.conf

# 暴露 NGINX 默认端口
EXPOSE 80

# 启动 NGINX
CMD ["nginx", "-g", "daemon off;"]
