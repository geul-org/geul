//ff:type feature=notation type=model
//ff:what codecURIEven, codecLanguageShort — Value codecs the notation names that core has no constant for
package notation

// codecURIEven is the even member of the URI byte-string pair (value/URI).
const codecURIEven uint8 = 0x44

// codecLanguageShort is LANGUAGE-SHORT-CODE (value/text/LANGUAGE-SHORT-CODE),
// one GEUL Language Registry number in 1 WORD.
const codecLanguageShort uint8 = 0x48
