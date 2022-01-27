package client

type operation interface {
	Process(state *albionState)
}

// Notes:
//   2020-08-31 (@phendryx): opAuctionGetItemsAverage removed from op codes
//			     based on public suggested changes and
//               @marleythemongolianmoose's findings:
//               "MarleyTheMongolianMoose: AuctionGetItemsAverage == 92 == kind
//               of looks like it disappears in the new one"

//OperationType used to identify operation types
//go:generate stringer -type=OperationType
type OperationType uint16

const (
	opUnused OperationType = iota
	opPing
	opJoin
	opCreateAccount
	opLogin
	opSendCrashLog
	opSendTraceRoute
	opSendVfxStats
	opSendGamePingInfo
	opCreateCharacter
	opDeleteCharacter
	opSelectCharacter
	opRedeemKeycode
	opGetGameServerByCluster
	opGetActiveSubscription
	opGetShopPurchaseUrl
	opGetBuyTrialDetails
	opGetReferralSeasonDetails
	opGetReferralLink
	opGetAvailableTrialKeys
	opGetShopTilesForCategory
	opMove
	opCastStart
	opCastCancel
	opTerminateToggleSpell
	opChannelingCancel
	opAttackBuildingStart
	opInventoryDestroyItem
	opInventoryMoveItem
	opInventoryRecoverItem
	opInventoryRecoverAllItems
	opInventorySplitStack
	opInventorySplitStackInto
	opGetClusterData
	opChangeCluster
	opConsoleCommand
	opChatMessage
	opReportClientError
	opRegisterToObject
	opUnRegisterFromObject
	opCraftBuildingChangeSettings
	opCraftBuildingTakeMoney
	opRepairBuildingChangeSettings
	opRepairBuildingTakeMoney
	opActionBuildingChangeSettings
	opHarvestStart
	opHarvestCancel
	opTakeSilver
	opActionOnBuildingStart
	opActionOnBuildingCancel
	opItemRerollQualityStart
	opItemRerollQualityCancel
	opInstallResourceStart
	opInstallResourceCancel
	opInstallSilver
	opBuildingFillNutrition
	opBuildingChangeRenovationState
	opBuildingBuySkin
	opBuildingClaim
	opBuildingGiveup
	opBuildingNutritionSilverStorageDeposit
	opBuildingNutritionSilverStorageWithdraw
	opBuildingNutritionSilverRewardSet
	opConstructionSiteCreate
	opPlaceableObjectPlace
	opPlaceableObjectPlaceCancel
	opPlaceableObjectPickup
	opFurnitureObjectUse
	opFarmableHarvest
	opFarmableFinishGrownItem
	opFarmableDestroy
	opFarmableGetProduct
	opTearDownConstructionSite
	opCastleGateUse
	opAuctionCreateRequest
	opAuctionCreateOffer
	opAuctionGetOffers
	opAuctionGetRequests
	opAuctionBuyOffer
	opAuctionAbortAuction
	opAuctionModifyAuction
	opAuctionAbortOffer
	opAuctionAbortRequest
	opAuctionSellRequest
	opAuctionGetFinishedAuctions
	opAuctionGetFinishedAuctionsCount
	opAuctionFetchAuction
	opAuctionGetMyOpenOffers
	opAuctionGetMyOpenRequests
	opAuctionGetMyOpenAuctions
	opAuctionGetItemAverageStats
	opAuctionGetItemAverageValue
	opContainerOpen
	opContainerClose
	opContainerManageSubContainer
	opRespawn
	opSuicide
	opJoinGuild
	opLeaveGuild
	opCreateGuild
	opInviteToGuild
	opDeclineGuildInvitation
	opKickFromGuild
	opDuellingChallengePlayer
	opDuellingAcceptChallenge
	opDuellingDenyChallenge
	opChangeClusterTax
	opClaimTerritory
	opGiveUpTerritory
	opChangeTerritoryAccessRights
	opGetMonolithInfo
	opGetClaimInfo
	opGetAttackInfo
	opGetTerritorySeasonPoints
	opGetAttackSchedule
	opScheduleAttack
	opGetMatches
	opGetMatchDetails
	opJoinMatch
	opLeaveMatch
	opChangeChatSettings
	opLogoutStart
	opLogoutCancel
	opClaimOrbStart
	opClaimOrbCancel
	opMatchLootChestOpeningStart
	opMatchLootChestOpeningCancel
	opDepositToGuildAccount
	opWithdrawalFromAccount
	opChangeGuildPayUpkeepFlag
	opChangeGuildTax
	opGetMyTerritories
	opMorganaCommand
	opGetServerInfo
	opInviteMercenaryToMatch
	opSubscribeToCluster
	opAnswerMercenaryInvitation
	opGetCharacterEquipment
	opGetCharacterSteamAchievements
	opGetCharacterStats
	opGetKillHistoryDetails
	opLearnMasteryLevel
	opReSpecAchievement
	opChangeAvatar
	opGetRankings
	opGetRank
	opGetGvgSeasonRankings
	opGetGvgSeasonRank
	opGetGvgSeasonHistoryRankings
	opGetGvgSeasonGuildMemberHistory
	opKickFromGvGMatch
	opGetChestLogs
	opGetAccessRightLogs
	opGetGuildAccountLogs
	opGetGuildAccountLogsLargeAmount
	opInviteToPlayerTrade
	opPlayerTradeCancel
	opPlayerTradeInvitationAccept
	opPlayerTradeAddItem
	opPlayerTradeRemoveItem
	opPlayerTradeAcceptTrade
	opPlayerTradeSetSilverOrGold
	opSendMiniMapPing
	opStuck
	opBuyRealEstate
	opClaimRealEstate
	opGiveUpRealEstate
	opChangeRealEstateOutline
	opGetMailInfos
	opGetMailCount
	opReadMail
	opSendNewMail
	opDeleteMail
	opMarkMailUnread
	opClaimAttachmentFromMail
	opUpdateLfgInfo
	opGetLfgInfos
	opGetMyGuildLfgInfo
	opGetLfgDescriptionText
	opLfgApplyToGuild
	opAnswerLfgGuildApplication
	opRegisterChatPeer
	opSendChatMessage
	opJoinChatChannel
	opLeaveChatChannel
	opSendWhisperMessage
	opSay
	opPlayEmote
	opStopEmote
	opGetClusterMapInfo
	opAccessRightsChangeSettings
	opMount
	opMountCancel
	opBuyJourney
	opSetSaleStatusForEstate
	opResolveGuildOrPlayerName
	opGetRespawnInfos
	opMakeHome
	opLeaveHome
	opResurrectionReply
	opAllianceCreate
	opAllianceDisband
	opAllianceGetMemberInfos
	opAllianceInvite
	opAllianceAnswerInvitation
	opAllianceCancelInvitation
	opAllianceKickGuild
	opAllianceLeave
	opAllianceChangeGoldPaymentFlag
	opAllianceGetDetailInfo
	opGetIslandInfos
	opAbandonMyIsland
	opBuyMyIsland
	opBuyGuildIsland
	opAbandonGuildIsland
	opUpgradeMyIsland
	opUpgradeGuildIsland
	opMoveMyIsland
	opMoveGuildIsland
	opTerritoryFillNutrition
	opTeleportBack
	opPartyInvitePlayer
	opPartyAnswerInvitation
	opPartyLeave
	opPartyKickPlayer
	opPartyMakeLeader
	opPartyChangeLootSetting
	opPartyMarkObject
	opPartySetRole
	opGetGuildMOTD
	opSetGuildMOTD
	opExitEnterStart
	opExitEnterCancel
	opQuestGiverRequest
	opGoldMarketGetBuyOffer
	opGoldMarketGetBuyOfferFromSilver
	opGoldMarketGetSellOffer
	opGoldMarketGetSellOfferFromSilver
	opGoldMarketBuyGold
	opGoldMarketSellGold
	opGoldMarketCreateSellOrder
	opGoldMarketCreateBuyOrder
	opGoldMarketGetInfos
	opGoldMarketCancelOrder
	opUnknown244
	opUnknown245
	opGoldMarketGetAverageInfo
	opSiegeCampClaimStart
	opSiegeCampClaimCancel
	opTreasureChestUsingStart
	opTreasureChestUsingCancel
	opUseLootChest
	opUseShrine
	opLaborerStartJob
	opLaborerTakeJobLoot
	opLaborerDismiss
	opLaborerMove
	opLaborerBuyItem
	opLaborerUpgrade
	opBuyPremium
	opBuyTrial
	opRealEstateGetAuctionData
	opRealEstateBidOnAuction
	opGetSiegeCampCooldown
	opFriendInvite
	opFriendAnswerInvitation
	opFriendCancelnvitation
	opFriendRemove
	opInventoryStack
	opInventorySort
	opEquipmentItemChangeSpell
	opExpeditionRegister
	opExpeditionRegisterCancel
	opJoinExpedition
	opDeclineExpeditionInvitation
	opVoteStart
	opVoteDoVote
	opRatingDoRate
	opEnteringExpeditionStart
	opEnteringExpeditionCancel
	opActivateExpeditionCheckPoint
	opArenaRegister
	opArenaRegisterCancel
	opArenaLeave
	opJoinArenaMatch
	opDeclineArenaInvitation
	opEnteringArenaStart
	opEnteringArenaCancel
	opArenaCustomMatch
	opArenaCustomMatchCreate
	opUpdateCharacterStatement
	opBoostFarmable
	opGetStrikeHistory
	opUseFunction
	opUsePortalEntrance
	opResetPortalBinding
	opQueryPortalBinding
	opClaimPaymentTransaction
	opChangeUseFlag
	opClientPerformanceStats
	opExtendedHardwareStats
	opClientLowMemoryWarning
	opTerritoryClaimStart
	opTerritoryClaimCancel
	opRequestAppStoreProducts
	opVerifyProductPurchase
	opQueryGuildPlayerStats
	opQueryAllianceGuildStats
	opTrackAchievements
	opSetAchievementsAutoLearn
	opDepositItemToGuildCurrency
	opWithdrawalItemFromGuildCurrency
	opAuctionSellSpecificItemRequest
	opFishingStart
	opFishingCasting
	opFishingCast
	opFishingCatch
	opFishingPull
	opFishingGiveLine
	opFishingFinish
	opFishingCancel
	opCreateGuildAccessTag
	opDeleteGuildAccessTag
	opRenameGuildAccessTag
	opFlagGuildAccessTagGuildPermission
	opAssignGuildAccessTag
	opRemoveGuildAccessTagFromPlayer
	opModifyGuildAccessTagEditors
	opRequestPublicAccessTags
	opChangeAccessTagPublicFlag
	opUpdateGuildAccessTag
	opSteamStartMicrotransaction
	opSteamFinishMicrotransaction
	opSteamIdHasActiveAccount
	opCheckEmailAccountState
	opLinkAccountToSteamId
	opBuyGvgSeasonBooster
	opChangeFlaggingPrepare
	opOverCharge
	opOverChargeEnd
	opRequestTrusted
	opChangeGuildLogo
	opPartyFinderRegisterForUpdates
	opPartyFinderUnregisterForUpdates
	opPartyFinderEnlistNewPartySearch
	opPartyFinderDeletePartySearch
	opPartyFinderChangePartySearch
	opPartyFinderChangeRole
	opPartyFinderApplyForGroup
	opPartyFinderAcceptOrDeclineApplyForGroup
	opPartyFinderGetEquipmentSnapshot
	opPartyFinderRegisterApplicants
	opPartyFinderUnregisterApplicants
	opPartyFinderFulltextSearch
	opPartyFinderRequestEquipmentSnapshot
	opGetPersonalSeasonTrackerData
	opUseConsumableFromInventory
	opClaimPersonalSeasonReward
	opEasyAntiCheatMessageToServer
	opSetNextTutorialState
	opAddPlayerToMuteList
	opRemovePlayerFromMuteList
	opProductShopUserEvent
	opGetVanityUnlocks
	opBuyVanityUnlocks
	opGetMountSkins
	opSetMountSkin
	opSetWardrobe
	opChangeCustomization
	opSetFavoriteIsland
	opGetGuildChallengePoints
	opTravelToHideout
	opSmartQueueJoin
	opSmartQueueLeave
	opSmartQueueSelectSpawnCluster
	opUpgradeHideout
	opInitHideoutAttackStart
	opInitHideoutAttackCancel
	opHideoutFillNutrition
	opHideoutGetInfo
	opHideoutGetOwnerInfo
	opHideoutSetTribute
	opOpenWorldAttackScheduleStart
	opOpenWorldAttackScheduleCancel
	opOpenWorldAttackConquerStart
	opOpenWorldAttackConquerCancel
	opGetOpenWorldAttackDetails
	opGetNextOpenWorldAttackScheduleTime
	opRecoverVaultFromHideout
	opGetGuildEnergyDrainInfo
	opChannelingUpdate
)
