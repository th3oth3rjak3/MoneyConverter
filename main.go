package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/th3oth3rjak3/MoneyConverter/money"
)

func main() {
	from := flag.String("from", "", "source currency code, required")
	to := flag.String("to", "", "target currency code, required")

	flag.Parse()

	value := flag.Arg(0)

	if value == "" {
		_, _ = fmt.Fprintln(os.Stderr, "Missing amount to convert")
		flag.Usage()
		os.Exit(1)
	}

	fromCurrency, err := money.ParseCurrency(*from)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse source currency %q: %s\n", *from, err.Error())
		os.Exit(1)
	}

	targetCurrency, err := money.ParseCurrency(*to)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse target currency %q: %s\n", *to, err.Error())
		os.Exit(1)
	}

	quantity, err := money.ParseDecimal(value)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse value %q: %s\n", value, err.Error())
		os.Exit(1)
	}

	fromAmount, err := money.NewAmount(quantity, fromCurrency)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	convertedAmount, err := money.Convert(fromAmount, targetCurrency)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to convert currency: %s", err.Error())
		os.Exit(1)
	}

	fmt.Printf("%s = %s\n", &fromAmount, &convertedAmount)
}
