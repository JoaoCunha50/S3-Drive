import { style } from '@vanilla-extract/css'
import { BlueSecondary, PurplePrimary, SurfaceGray } from '../../styles/colors'
import { vars } from '../../styles/theme'

export const wrapper = style({
    minHeight: '100dvh',
    display: 'grid',
    placeItems: 'center',
    padding: '0px 24px 0px 24px',
})

export const paper = style({
    width: '45%',
    maxWidth: '50%',
    backgroundColor: SurfaceGray,
    padding: '5% 6% 5% 6%',
    borderRadius: vars.radiusDefault,
})

export const header = style({
    display: 'flex',
    flexDirection: 'row',
    marginBottom: '35px',
    textAlign: 'center',
    justifyContent: 'center',
    alignItems: 'center',
    marginRight: '5px',
    gap: '5px',
})

export const title = style({
    margin: 0,
    fontWeight: 700,
    fontSize: '40px',
    background: `linear-gradient(to right, ${PurplePrimary}, ${BlueSecondary})`,
    WebkitBackgroundClip: 'text',
    color: 'transparent',
    lineHeight: 1.2,
})

export const form = style({
    display: 'grid',
    gap: '12px',
    marginTop: '0px',
})

export const actions = style({
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginTop: '20px',
})

export const error = style({
    color: 'red',
    fontSize: '10px',
})
