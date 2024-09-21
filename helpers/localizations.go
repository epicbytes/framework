package helpers

import (
	"context"
	"github.com/goodsign/monday"
	"github.com/samber/lo"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"time"
)

const (
	time_simple_pm  = "3:04PM"
	time_complex_pm = "3:04:05 PM"

	time_full     = "15:04"
	time_full_sec = "15:04:05"

	date_d_m_y  = "01/02/2006"
	date_m_d_y  = "02/01/2006"
	date_dotted = "02.01.2006"
	date_kebab  = "02-01-2006"
	date_ch     = "2006年01月02日"
)

func getLangBaseFromCtx(ctx context.Context) language.Tag {
	if lng, ok := ctx.Value("currentLanguageTag").(language.Tag); ok {
		return lng
	}
	return language.English
}

func CurrencyFormatter(ctx context.Context, amount interface{}) string {
	currencyPrinter := message.NewPrinter(getLangBaseFromCtx(ctx))
	return currencyPrinter.Sprint(currency.Symbol(currency.RUB.Amount(amount)))
}

// Date internationalization's

func GetLocalizedDateFormat(locale language.Tag) string {
	switch locale {
	case language.AmericanEnglish:
		return date_d_m_y
	case language.French, language.CanadianFrench:
		return date_kebab
	case language.Chinese:
		return date_ch
	case language.Russian:
		return date_dotted
	default:
		return date_m_d_y
	}
}

func GetLocalizedTimeFormat(locale language.Tag, isSimple bool) string {
	switch locale {
	case language.CanadianFrench, language.English:
		return lo.Ternary(isSimple, time_simple_pm, time_complex_pm)
	default:
		return lo.Ternary(isSimple, time_full, time_full_sec)
	}
}

func GetLocalizedDateTimeFormat(locale language.Tag, isSimple bool) string {
	return GetLocalizedDateFormat(locale) + " " + GetLocalizedTimeFormat(locale, isSimple)
}

func GetLocalizedDateTimeFormatFromCtx(ctx context.Context, isSimple bool) string {
	return GetLocalizedDateTimeFormat(getLangBaseFromCtx(ctx), isSimple)
}

func GetDateFormatted(ctx context.Context, data int64) string {
	td := time.Unix(data, 0)
	//translation := monday.Format(t, layout, monday.LocaleBgBG) // Instead of t.Format(layout)
	translation := monday.Format(td, GetLocalizedDateTimeFormatFromCtx(ctx, true), monday.Locale(getLangBaseFromCtx(ctx).String())) // Instead of t.Format(layout)
	return translation
}
