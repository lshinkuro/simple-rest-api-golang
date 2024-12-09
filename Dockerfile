# Gunakan base image resmi Golang
FROM golang:1.20-alpine

# Atur direktori kerja di dalam container
WORKDIR /app

# Copy semua file dari project lokal ke dalam container
COPY . .

# Unduh semua dependency yang dibutuhkan
RUN go mod tidy

# Build aplikasi Go menjadi binary file
RUN go build -o main .

# Ekspos port (ubah sesuai kebutuhan aplikasi Anda)
EXPOSE 8080

# Jalankan aplikasi saat container dimulai
CMD ["./main"]
