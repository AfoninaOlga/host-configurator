TARGET=configurator

.PHONY: all
all: $(TARGET)

.PHONY: $(TARGET)
$(TARGET): tidy
	go build -o $@ main.go

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: clean
clean:
	$(RM) $(TARGET)
