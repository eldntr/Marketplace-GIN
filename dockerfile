# Gunakan image resmi MySQL
FROM mysql:8

# Set environment variable untuk root password dan user custom
ENV MYSQL_ROOT_PASSWORD=my-root-password
ENV MYSQL_DATABASE=mydb
ENV MYSQL_USER=myuser
ENV MYSQL_PASSWORD=mypassword

# Salin file konfigurasi jika perlu (opsional)
# COPY ./my_custom_config.cnf /etc/mysql/conf.d/

# Expose port 3306 untuk akses luar
EXPOSE 3306

# MySQL akan secara otomatis memulai dan siap digunakan