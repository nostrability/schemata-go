package schemata

import (
	"embed"
	"encoding/json"
	"sort"
	"sync"
)

//go:embed schemas
var schemasFS embed.FS

var registry = map[string]string{
	"kind0ContentSchema": "schemas/nips/nip-01/kind-0/schema.content.json",
	"kind0Schema": "schemas/nips/nip-01/kind-0/schema.json",
	"kind1Schema": "schemas/nips/nip-01/kind-1/schema.json",
	"clientCloseSchema": "schemas/nips/nip-01/messages/client-close/schema.json",
	"clientEventSchema": "schemas/nips/nip-01/messages/client-event/schema.json",
	"clientReqSchema": "schemas/nips/nip-01/messages/client-req/schema.json",
	"filterSchema": "schemas/nips/nip-01/messages/filter/schema.json",
	"relayClosedSchema": "schemas/nips/nip-01/messages/relay-closed/schema.json",
	"relayEoseSchema": "schemas/nips/nip-01/messages/relay-eose/schema.json",
	"relayEventSchema": "schemas/nips/nip-01/messages/relay-event/schema.json",
	"relayNoticeSchema": "schemas/nips/nip-01/messages/relay-notice/schema.json",
	"relayOkSchema": "schemas/nips/nip-01/messages/relay-ok/schema.json",
	"noteSchema": "schemas/nips/nip-01/note/schema.json",
	"noteUnsignedSchema": "schemas/nips/nip-01/note-unsigned/schema.json",
	"secp256k1Schema": "schemas/nips/nip-01/secp256k1/schema.json",
	"aTagSchema": "schemas/nips/nip-01/tag/a/schema.json",
	"dTagSchema": "schemas/nips/nip-01/tag/d/schema.json",
	"eTagSchema": "schemas/nips/nip-01/tag/e/schema.json",
	"pTagSchema": "schemas/nips/nip-01/tag/p/schema.json",
	"tagSchema": "schemas/nips/nip-01/tag/schema.json",
	"tTagSchema": "schemas/nips/nip-01/tag/t/schema.json",
	"kind3ContentSchema": "schemas/nips/nip-02/kind-3/schema.content.json",
	"kind3Schema": "schemas/nips/nip-02/kind-3/schema.json",
	"kind4Schema": "schemas/nips/nip-04/kind-4/schema.json",
	"identifierSchema": "schemas/nips/nip-05/identifier/schema.json",
	"wellKnownSchema": "schemas/nips/nip-05/well-known/schema.json",
	"kind5Schema": "schemas/nips/nip-09/kind-5/schema.json",
	"nip11Schema": "schemas/nips/nip-11/schema.json",
	"kind10050Schema": "schemas/nips/nip-17/kind-10050/schema.json",
	"kind14Schema": "schemas/nips/nip-17/kind-14/schema.json",
	"kind15Schema": "schemas/nips/nip-17/kind-15/schema.json",
	"blurhashTagSchema": "schemas/nips/nip-17/tag/blurhash/schema.json",
	"decryptionkeyTagSchema": "schemas/nips/nip-17/tag/decryption-key/schema.json",
	"decryptionnonceTagSchema": "schemas/nips/nip-17/tag/decryption-nonce/schema.json",
	"dimTagSchema": "schemas/nips/nip-17/tag/dim/schema.json",
	"encryptionalgorithmTagSchema": "schemas/nips/nip-17/tag/encryption-algorithm/schema.json",
	"fallbackTagSchema": "schemas/nips/nip-17/tag/fallback/schema.json",
	"filetypeTagSchema": "schemas/nips/nip-17/tag/file-type/schema.json",
	"oxTagSchema": "schemas/nips/nip-17/tag/ox/schema.json",
	"relayTagSchema": "schemas/nips/nip-17/tag/relay/schema.json",
	"sizeTagSchema": "schemas/nips/nip-17/tag/size/schema.json",
	"subjectTagSchema": "schemas/nips/nip-17/tag/subject/schema.json",
	"thumbTagSchema": "schemas/nips/nip-17/tag/thumb/schema.json",
	"xTagSchema": "schemas/nips/nip-17/tag/x/schema.json",
	"kind16Schema": "schemas/nips/nip-18/kind-16/schema.json",
	"kind6Schema": "schemas/nips/nip-18/kind-6/schema.json",
	"kTagSchema": "schemas/nips/nip-18/tag/k/schema.json",
	"qTagSchema": "schemas/nips/nip-18/tag/q/schema.json",
	"kind1111Schema": "schemas/nips/nip-22/kind-1111/schema.json",
	"ATagSchema": "schemas/nips/nip-22/tag/_A/schema.json",
	"ETagSchema": "schemas/nips/nip-22/tag/_E/schema.json",
	"KTagSchema": "schemas/nips/nip-22/tag/_K/schema.json",
	"PTagSchema": "schemas/nips/nip-22/tag/_P/schema.json",
	"kind30023Schema": "schemas/nips/nip-23/kind-30023/schema.json",
	"kind30024Schema": "schemas/nips/nip-23/kind-30024/schema.json",
	"publishedatTagSchema": "schemas/nips/nip-23/tag/published_at/schema.json",
	"kind17Schema": "schemas/nips/nip-25/kind-17/schema.json",
	"kind7Schema": "schemas/nips/nip-25/kind-7/schema.json",
	"ereactTagSchema": "schemas/nips/nip-25/tag/e-react/schema.json",
	"emojiTagSchema": "schemas/nips/nip-25/tag/emoji/schema.json",
	"kind1985Schema": "schemas/nips/nip-32/kind-1985/schema.json",
	"LTagSchema": "schemas/nips/nip-32/tag/_L/schema.json",
	"lTagSchema": "schemas/nips/nip-32/tag/l/schema.json",
	"kind1617Schema": "schemas/nips/nip-34/kind-1617/schema.json",
	"kind1621Schema": "schemas/nips/nip-34/kind-1621/schema.json",
	"kind1630Schema": "schemas/nips/nip-34/kind-1630/schema.json",
	"kind1631Schema": "schemas/nips/nip-34/kind-1631/schema.json",
	"kind1632Schema": "schemas/nips/nip-34/kind-1632/schema.json",
	"kind1633Schema": "schemas/nips/nip-34/kind-1633/schema.json",
	"kind30617Schema": "schemas/nips/nip-34/kind-30617/schema.json",
	"kind30618Schema": "schemas/nips/nip-34/kind-30618/schema.json",
	"appliedascommitsTagSchema": "schemas/nips/nip-34/tag/applied-as-commits/schema.json",
	"cloneTagSchema": "schemas/nips/nip-34/tag/clone/schema.json",
	"commitTagSchema": "schemas/nips/nip-34/tag/commit/schema.json",
	"commitpgpsigTagSchema": "schemas/nips/nip-34/tag/commit-pgp-sig/schema.json",
	"committerTagSchema": "schemas/nips/nip-34/tag/committer/schema.json",
	"erootTagSchema": "schemas/nips/nip-34/tag/e-root/schema.json",
	"estatusreplyTagSchema": "schemas/nips/nip-34/tag/e-status-reply/schema.json",
	"headTagSchema": "schemas/nips/nip-34/tag/head/schema.json",
	"maintainersTagSchema": "schemas/nips/nip-34/tag/maintainers/schema.json",
	"mergecommitTagSchema": "schemas/nips/nip-34/tag/merge-commit/schema.json",
	"nameTagSchema": "schemas/nips/nip-34/tag/name/schema.json",
	"parentcommitTagSchema": "schemas/nips/nip-34/tag/parent-commit/schema.json",
	"reucTagSchema": "schemas/nips/nip-34/tag/r-euc/schema.json",
	"refTagSchema": "schemas/nips/nip-34/tag/ref/schema.json",
	"webTagSchema": "schemas/nips/nip-34/tag/web/schema.json",
	"contentwarningTagSchema": "schemas/nips/nip-36/tag/content-warning/schema.json",
	"kind30315Schema": "schemas/nips/nip-38/kind-30315/schema.json",
	"expirationTagSchema": "schemas/nips/nip-38/tag/expiration/schema.json",
	"statustypeTagSchema": "schemas/nips/nip-38/tag/status-type/schema.json",
	"clientAuthSchema": "schemas/nips/nip-40/messages/client-auth/schema.json",
	"relayAuthSchema": "schemas/nips/nip-40/messages/relay-auth/schema.json",
	"kind22242Schema": "schemas/nips/nip-42/kind-22242/schema.json",
	"challengeTagSchema": "schemas/nips/nip-42/tag/challenge/schema.json",
	"kind31922Schema": "schemas/nips/nip-52/kind-31922/schema.json",
	"kind31923Schema": "schemas/nips/nip-52/kind-31923/schema.json",
	"kind31924Schema": "schemas/nips/nip-52/kind-31924/schema.json",
	"kind31925Schema": "schemas/nips/nip-52/kind-31925/schema.json",
	"endtzidTagSchema": "schemas/nips/nip-52/tag/end_tzid/schema.json",
	"fbTagSchema": "schemas/nips/nip-52/tag/fb/schema.json",
	"gTagSchema": "schemas/nips/nip-52/tag/g/schema.json",
	"locationTagSchema": "schemas/nips/nip-52/tag/location/schema.json",
	"starttzidTagSchema": "schemas/nips/nip-52/tag/start_tzid/schema.json",
	"statusTagSchema": "schemas/nips/nip-52/tag/status/schema.json",
	"kind10312Schema": "schemas/nips/nip-53/kind-10312/schema.json",
	"kind1311Schema": "schemas/nips/nip-53/kind-1311/schema.json",
	"kind30311Schema": "schemas/nips/nip-53/kind-30311/schema.json",
	"kind30312Schema": "schemas/nips/nip-53/kind-30312/schema.json",
	"kind30313Schema": "schemas/nips/nip-53/kind-30313/schema.json",
	"aliveTagSchema": "schemas/nips/nip-53/tag/a-live/schema.json",
	"aroomTagSchema": "schemas/nips/nip-53/tag/a-room/schema.json",
	"currentparticipantsTagSchema": "schemas/nips/nip-53/tag/current_participants/schema.json",
	"endpointTagSchema": "schemas/nips/nip-53/tag/endpoint/schema.json",
	"endsTagSchema": "schemas/nips/nip-53/tag/ends/schema.json",
	"handTagSchema": "schemas/nips/nip-53/tag/hand/schema.json",
	"imageTagSchema": "schemas/nips/nip-53/tag/image/schema.json",
	"pinnedTagSchema": "schemas/nips/nip-53/tag/pinned/schema.json",
	"recordingTagSchema": "schemas/nips/nip-53/tag/recording/schema.json",
	"relaysTagSchema": "schemas/nips/nip-53/tag/relays/schema.json",
	"roomTagSchema": "schemas/nips/nip-53/tag/room/schema.json",
	"serviceTagSchema": "schemas/nips/nip-53/tag/service/schema.json",
	"startsTagSchema": "schemas/nips/nip-53/tag/starts/schema.json",
	"statusliveTagSchema": "schemas/nips/nip-53/tag/status-live/schema.json",
	"statusroomTagSchema": "schemas/nips/nip-53/tag/status-room/schema.json",
	"streamingTagSchema": "schemas/nips/nip-53/tag/streaming/schema.json",
	"summaryTagSchema": "schemas/nips/nip-53/tag/summary/schema.json",
	"titleTagSchema": "schemas/nips/nip-53/tag/title/schema.json",
	"totalparticipantsTagSchema": "schemas/nips/nip-53/tag/total_participants/schema.json",
	"kind1984Schema": "schemas/nips/nip-56/kind-1984/schema.json",
	"serverTagSchema": "schemas/nips/nip-56/tag/server/schema.json",
	"kind9734Schema": "schemas/nips/nip-57/kind-9734/schema.json",
	"kind9735Schema": "schemas/nips/nip-57/kind-9735/schema.json",
	"amountTagSchema": "schemas/nips/nip-57/tag/amount/schema.json",
	"bolt11TagSchema": "schemas/nips/nip-57/tag/bolt11/schema.json",
	"descriptionTagSchema": "schemas/nips/nip-57/tag/description/schema.json",
	"lnurlTagSchema": "schemas/nips/nip-57/tag/lnurl/schema.json",
	"preimageTagSchema": "schemas/nips/nip-57/tag/preimage/schema.json",
	"kind1059Schema": "schemas/nips/nip-59/kind-1059/schema.json",
	"kind13Schema": "schemas/nips/nip-59/kind-13/schema.json",
	"kind10019Schema": "schemas/nips/nip-61/kind-10019/schema.json",
	"kind7376Schema": "schemas/nips/nip-61/kind-7376/schema.json",
	"kind9321Schema": "schemas/nips/nip-61/kind-9321/schema.json",
	"eredeemedTagSchema": "schemas/nips/nip-61/tag/e-redeemed/schema.json",
	"mintTagSchema": "schemas/nips/nip-61/tag/mint/schema.json",
	"proofTagSchema": "schemas/nips/nip-61/tag/proof/schema.json",
	"pubkeyTagSchema": "schemas/nips/nip-61/tag/pubkey/schema.json",
	"uTagSchema": "schemas/nips/nip-61/tag/u/schema.json",
	"kind10002Schema": "schemas/nips/nip-65/kind-10002/schema.json",
	"rTagSchema": "schemas/nips/nip-65/tag/r/schema.json",
	"kind20Schema": "schemas/nips/nip-68/kind-20/schema.json",
	"imetaTagSchema": "schemas/nips/nip-68/tag/imeta/schema.json",
	"iTagSchema": "schemas/nips/nip-73/tag/i/schema.json",
	"kind30078Schema": "schemas/nips/nip-78/kind-30078/schema.json",
	"kind9802Schema": "schemas/nips/nip-84/kind-9802/schema.json",
	"commentTagSchema": "schemas/nips/nip-84/tag/comment/schema.json",
	"contextTagSchema": "schemas/nips/nip-84/tag/context/schema.json",
	"kind1018Schema": "schemas/nips/nip-88/kind-1018/schema.json",
	"kind1068Schema": "schemas/nips/nip-88/kind-1068/schema.json",
	"endsAtTagSchema": "schemas/nips/nip-88/tag/endsAt/schema.json",
	"optionTagSchema": "schemas/nips/nip-88/tag/option/schema.json",
	"polltypeTagSchema": "schemas/nips/nip-88/tag/polltype/schema.json",
	"responseTagSchema": "schemas/nips/nip-88/tag/response/schema.json",
	"kind31989Schema": "schemas/nips/nip-89/kind-31989/schema.json",
	"kind31990Schema": "schemas/nips/nip-89/kind-31990/schema.json",
	"clientTagSchema": "schemas/nips/nip-89/tag/client/schema.json",
	"kind30402Schema": "schemas/nips/nip-99/kind-30402/schema.json",
	"kind30403Schema": "schemas/nips/nip-99/kind-30403/schema.json",
	"kind10063Schema": "schemas/nips/nip-b7/kind-10063/schema.json",
	"kind10051Schema": "schemas/mips/mip-00/kind-10051/schema.json",
	"kind443Schema": "schemas/mips/mip-00/kind-443/schema.json",
	"encodingTagSchema": "schemas/mips/mip-00/tag/encoding/schema.json",
	"mlsciphersuiteTagSchema": "schemas/mips/mip-00/tag/mls_ciphersuite/schema.json",
	"mlsextensionsTagSchema": "schemas/mips/mip-00/tag/mls_extensions/schema.json",
	"mlsprotocolversionTagSchema": "schemas/mips/mip-00/tag/mls_protocol_version/schema.json",
	"kind444Schema": "schemas/mips/mip-02/kind-444/schema.json",
	"kind445Schema": "schemas/mips/mip-03/kind-445/schema.json",
	"hTagSchema": "schemas/mips/mip-03/tag/h/schema.json",
	"messagefilterSchema": "schemas/@/message/filter.json",
	"nip05Schema": "schemas/@/nip05.json",
	"nostrwellknownSchema": "schemas/@/nostr-well-known.json",
	"notekind22242Schema": "schemas/@/note/kind-22242.json",
	"noteunsignedSchema": "schemas/@/note-unsigned.json",
	"authrelayTagSchema": "schemas/@/tag/auth-relay.json",
	"dmrelayTagSchema": "schemas/@/tag/dm-relay.json",
	"externalcontentiTagSchema": "schemas/@/tag/external-content-i.json",
	"externalcontentkTagSchema": "schemas/@/tag/external-content-k.json",
	"gitrefTagSchema": "schemas/@/tag/git-ref.json",
	"mip00iTagSchema": "schemas/@/tag/mip-00_i.json",
	"mlsclientTagSchema": "schemas/@/tag/mls_client.json",
	"reactioneTagSchema": "schemas/@/tag/reaction-e.json",
}

var (
	cache     map[string]*json.RawMessage
	cacheOnce sync.Once
)

func loadCache() {
	cache = make(map[string]*json.RawMessage, len(registry))
	for key, path := range registry {
		data, err := schemasFS.ReadFile(path)
		if err != nil {
			continue
		}
		raw := json.RawMessage(data)
		cache[key] = &raw
	}
}

// Get returns the schema JSON for the given registry key.
func Get(key string) (json.RawMessage, bool) {
	cacheOnce.Do(loadCache)
	raw, ok := cache[key]
	if !ok {
		return nil, false
	}
	return *raw, true
}

// Keys returns all available schema registry keys, sorted alphabetically.
func Keys() []string {
	cacheOnce.Do(loadCache)
	keys := make([]string, 0, len(cache))
	for k := range cache {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
