#!/bin/sh

# Coverage script for CI pipeline
echo "Processing coverage results..."

# Check if coverage.out exists
if [ -f "coverage.out" ]; then
    echo "Coverage file found: coverage.out"
    
    # Generate HTML coverage report
    go tool cover -html=coverage.out -o coverage.html
    
    # Show coverage summary
    go tool cover -func=coverage.out
    
    echo "Coverage processing completed"
else
    echo "Warning: coverage.out not found"
    exit 1
fi 